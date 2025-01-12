// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Types

import "strconv"

type ShockerCommandType byte

const (
	ShockerCommandTypeStop    ShockerCommandType = 0
	ShockerCommandTypeShock   ShockerCommandType = 1
	ShockerCommandTypeVibrate ShockerCommandType = 2
	ShockerCommandTypeSound   ShockerCommandType = 3
)

var EnumNamesShockerCommandType = map[ShockerCommandType]string{
	ShockerCommandTypeStop:    "Stop",
	ShockerCommandTypeShock:   "Shock",
	ShockerCommandTypeVibrate: "Vibrate",
	ShockerCommandTypeSound:   "Sound",
}

var EnumValuesShockerCommandType = map[string]ShockerCommandType{
	"Stop":    ShockerCommandTypeStop,
	"Shock":   ShockerCommandTypeShock,
	"Vibrate": ShockerCommandTypeVibrate,
	"Sound":   ShockerCommandTypeSound,
}

func (v ShockerCommandType) String() string {
	if s, ok := EnumNamesShockerCommandType[v]; ok {
		return s
	}
	return "ShockerCommandType(" + strconv.FormatInt(int64(v), 10) + ")"
}
