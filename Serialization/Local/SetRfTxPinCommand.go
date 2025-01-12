// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Local

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SetRfTxPinCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsSetRfTxPinCommand(buf []byte, offset flatbuffers.UOffsetT) *SetRfTxPinCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SetRfTxPinCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishSetRfTxPinCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSetRfTxPinCommand(buf []byte, offset flatbuffers.UOffsetT) *SetRfTxPinCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SetRfTxPinCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSetRfTxPinCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SetRfTxPinCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SetRfTxPinCommand) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SetRfTxPinCommand) Pin() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SetRfTxPinCommand) MutatePin(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func SetRfTxPinCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SetRfTxPinCommandAddPin(builder *flatbuffers.Builder, pin int8) {
	builder.PrependInt8Slot(0, pin, 0)
}
func SetRfTxPinCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
