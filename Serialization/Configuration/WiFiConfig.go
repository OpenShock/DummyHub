// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Configuration

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WiFiConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsWiFiConfig(buf []byte, offset flatbuffers.UOffsetT) *WiFiConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WiFiConfig{}
	x.Init(buf, n+offset)
	return x
}

func FinishWiFiConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsWiFiConfig(buf []byte, offset flatbuffers.UOffsetT) *WiFiConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WiFiConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedWiFiConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *WiFiConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WiFiConfig) Table() flatbuffers.Table {
	return rcv._tab
}

/// Access point SSID
func (rcv *WiFiConfig) ApSsid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Access point SSID
/// Hub hostname
func (rcv *WiFiConfig) Hostname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Hub hostname
/// WiFi network credentials
func (rcv *WiFiConfig) Credentials(obj *WiFiCredentials, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *WiFiConfig) CredentialsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// WiFi network credentials
func WiFiConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func WiFiConfigAddApSsid(builder *flatbuffers.Builder, apSsid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(apSsid), 0)
}
func WiFiConfigAddHostname(builder *flatbuffers.Builder, hostname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(hostname), 0)
}
func WiFiConfigAddCredentials(builder *flatbuffers.Builder, credentials flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(credentials), 0)
}
func WiFiConfigStartCredentialsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func WiFiConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
