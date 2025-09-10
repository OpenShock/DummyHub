package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/gorilla/websocket"

	Gateway "OpenShock/Serialization/Gateway"
	Types "OpenShock/Serialization/Types"
)

// Global variable to track the program's start time
var startTime time.Time

type Config struct {
	AuthToken string `json:"auth_token"`
	BaseURL   string `json:"base_url"`
}

const (
	configFile = "config.json"
)

func initializeUptime() {
	startTime = time.Now()
}

func getUptimeMs() uint64 {
	return uint64(time.Since(startTime).Milliseconds())
}

func loadConfig() (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			var newConfig Config = Config{BaseURL: "https://api.openshock.app"}
			saveConfig(&newConfig)
			return &newConfig, nil
		}
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	if config.BaseURL == "" {
		config.BaseURL = "https://api.openshock.app"
	}
	return &config, nil
}

func saveConfig(config *Config) error {
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

func getAuthToken(baseURL, pairingCode string) (string, error) {
	url := fmt.Sprintf("%s/1/device/pair/%s", baseURL, pairingCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to retrieve token: %s", resp.Status)
	}

	// Try decoding as JSON object first
	var jsonResponse struct {
		Token string `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err == nil && jsonResponse.Token != "" {
		return jsonResponse.Token, nil
	}

	// If JSON object decoding fails, reset the body and try decoding as plain string
	resp.Body.Close()
	resp, err = client.Get(url) // Re-fetch the response to reset the body
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var token string
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return token, nil
}

func getLCGURL(baseURL, authToken string) (string, error) {
	url := fmt.Sprintf("%s/1/device/assignLCG", baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("DeviceToken", authToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to retrieve LCG URL: %s", resp.Status)
	}

	var response struct {
		Data struct {
			FQDN    string `json:"fqdn"`
			Country string `json:"country"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}
	return response.Data.FQDN, nil
}

func buildSemVer(builder *flatbuffers.Builder, major, minor, patch uint16, build, prerelease *string) flatbuffers.UOffsetT {
	buildOffset := flatbuffers.UOffsetT(0)
	prereleaseOffset := flatbuffers.UOffsetT(0)

	if build != nil {
		buildOffset = builder.CreateString(*build)
	}
	if prerelease != nil {
		prereleaseOffset = builder.CreateString(*prerelease)
	}

	Types.SemVerStart(builder)
	Types.SemVerAddMajor(builder, major)
	Types.SemVerAddMinor(builder, minor)
	Types.SemVerAddPatch(builder, patch)
	if build != nil {
		Types.SemVerAddBuild(builder, buildOffset)
	}
	if prerelease != nil {
		Types.SemVerAddPrerelease(builder, prereleaseOffset)
	}
	return Types.SemVerEnd(builder)
}

func buildBootStatus(builder *flatbuffers.Builder, bootType Types.FirmwareBootType, firmwareVersion flatbuffers.UOffsetT, otaUpdateID int) flatbuffers.UOffsetT {
	Gateway.BootStatusStart(builder)
	Gateway.BootStatusAddBootType(builder, bootType)
	Gateway.BootStatusAddFirmwareVersion(builder, firmwareVersion)
	Gateway.BootStatusAddOtaUpdateId(builder, int32(otaUpdateID))
	return Gateway.BootStatusEnd(builder)
}

func buildPong(builder *flatbuffers.Builder, uptime uint64, rssi int32) flatbuffers.UOffsetT {
	Gateway.PongStart(builder)
	Gateway.PongAddUptime(builder, uptime)
	Gateway.PongAddRssi(builder, rssi)
	return Gateway.PongEnd(builder)
}

func buildHubToGatewayMessage(builder *flatbuffers.Builder, payloadType Gateway.HubToGatewayMessagePayload, payload flatbuffers.UOffsetT) []byte {
	Gateway.HubToGatewayMessageStart(builder)
	Gateway.HubToGatewayMessageAddPayloadType(builder, payloadType)
	Gateway.HubToGatewayMessageAddPayload(builder, payload)
	message := Gateway.HubToGatewayMessageEnd(builder)
	builder.Finish(message)
	return builder.FinishedBytes()
}

func sendBootStatus(conn *websocket.Conn) error {
	builder := flatbuffers.NewBuilder(1024)

	semver := buildSemVer(builder, 1, 0, 0, nil, nil) // Example: Major=1, Minor=0, Patch=0, no build or prerelease
	bootStatus := buildBootStatus(builder, Types.FirmwareBootTypeNormal, semver, 0)
	message := buildHubToGatewayMessage(builder, Gateway.HubToGatewayMessagePayloadBootStatus, bootStatus)

	if err := conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return fmt.Errorf("failed to send boot status message: %w", err)
	}

	return nil
}

func sendPong(conn *websocket.Conn) error {
	builder := flatbuffers.NewBuilder(1024)

	pong := buildPong(builder, getUptimeMs(), -65)
	message := buildHubToGatewayMessage(builder, Gateway.HubToGatewayMessagePayloadPong, pong)

	if err := conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return fmt.Errorf("failed to send boot status message: %w", err)
	}

	return nil
}

func connectWebSocket(authToken, lcgURL string) error {
	headers := http.Header{}
	headers.Add("Firmware-Version", "1.0.0-Dummy")
	headers.Add("Device-Token", authToken)

	fullURL := fmt.Sprintf("wss://%s/2/ws/hub", lcgURL)
	log.Printf("Connecting to WebSocket at %s", fullURL)

	conn, _, err := websocket.DefaultDialer.Dial(fullURL, headers)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	defer conn.Close()
	log.Println("Connected to WebSocket")

	err = sendBootStatus(conn)
	if err != nil {
		return err
	}

	return handleWebSocketMessages(conn)
}

// Handle Incoming WebSocket Messages
func handleWebSocketMessages(conn *websocket.Conn) error {
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("WebSocket read error: %w", err)
		}

		message := Gateway.GetRootAsGatewayToHubMessage(bytes, 0)
		switch message.PayloadType() {
		case Gateway.GatewayToHubMessagePayloadPing:
			err := sendPong(conn)
			if err != nil {
				return err
			}
		case Gateway.GatewayToHubMessagePayloadTrigger:
			err := handleTrigger(message)
			if err != nil {
				return err
			}
		case Gateway.GatewayToHubMessagePayloadShockerCommandList:
			err := handleCommandList(message)
			if err != nil {
				return err
			}
		default:
			log.Printf("Unknown message type: %d", message.PayloadType())
		}
	}
}

func handleTrigger(message *Gateway.GatewayToHubMessage) error {
	table := flatbuffers.Table{}
	if !message.Payload(&table) {
		return fmt.Errorf("failed to deserialize commandlist")
	}
	trigger := Gateway.Trigger{}
	trigger.Init(table.Bytes, table.Pos)

	switch trigger.Type() {
	case Gateway.TriggerTypeRestart:
		log.Println("[TRIGGER] Received restart command!")
	case Gateway.TriggerTypeEmergencyStop:
		log.Println("[TRIGGER] Emergency stop activated!")
	case Gateway.TriggerTypeCaptivePortalEnable:
		log.Println("[TRIGGER] Captive portal enabled!")
	case Gateway.TriggerTypeCaptivePortalDisable:
		log.Println("[TRIGGER] Captive portal disabled!")
	default:
		log.Printf("Unknown message type: %d", message.PayloadType())
	}

	return nil
}

func handleCommandList(message *Gateway.GatewayToHubMessage) error {
	table := flatbuffers.Table{}
	if !message.Payload(&table) {
		return fmt.Errorf("failed to deserialize commandlist")
	}
	commandlist := Gateway.ShockerCommandList{}
	commandlist.Init(table.Bytes, table.Pos)

	ncommands := commandlist.CommandsLength()

	command := Gateway.ShockerCommand{}
	for i := 0; i < ncommands; i++ {
		if !commandlist.Commands(&command, i) {
			return fmt.Errorf("failed to deserialize command")
		}

		fmt.Printf("[COMMAND] %s at %v%% for %v seconds\n", command.Type().String(), command.Intensity(), float32(command.Duration())/1000.0)
	}

	return nil
}

func main() {
	initializeUptime()

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if config.AuthToken == "" {
		fmt.Print("Enter pairing code: ")
		reader := bufio.NewReader(os.Stdin)
		pairingCode, _ := reader.ReadString('\n')
		pairingCode = strings.TrimSpace(pairingCode)

		config.AuthToken, err = getAuthToken(config.BaseURL, pairingCode)
		if err != nil {
			log.Fatalf("Failed to get auth token: %v", err)
		}

		if err := saveConfig(config); err != nil {
			log.Fatalf("Failed to save config: %v", err)
		}
	}

	lcgURL, err := getLCGURL(config.BaseURL, config.AuthToken)
	if err != nil {
		log.Fatalf("Failed to get LCG URL: %v", err)
	}

	log.Printf("Got lcgURL: %s", lcgURL)

	if err := connectWebSocket(config.AuthToken, lcgURL); err != nil {
		log.Fatalf("WebSocket error: %v", err)
	}
}
