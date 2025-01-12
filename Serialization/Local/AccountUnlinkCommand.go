// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Local

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AccountUnlinkCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsAccountUnlinkCommand(buf []byte, offset flatbuffers.UOffsetT) *AccountUnlinkCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AccountUnlinkCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishAccountUnlinkCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsAccountUnlinkCommand(buf []byte, offset flatbuffers.UOffsetT) *AccountUnlinkCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AccountUnlinkCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedAccountUnlinkCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *AccountUnlinkCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccountUnlinkCommand) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AccountUnlinkCommand) Placeholder() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AccountUnlinkCommand) MutatePlaceholder(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func AccountUnlinkCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AccountUnlinkCommandAddPlaceholder(builder *flatbuffers.Builder, placeholder bool) {
	builder.PrependBoolSlot(0, placeholder, false)
}
func AccountUnlinkCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
