// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Local

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OtaUpdateStartUpdateCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsOtaUpdateStartUpdateCommand(buf []byte, offset flatbuffers.UOffsetT) *OtaUpdateStartUpdateCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OtaUpdateStartUpdateCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishOtaUpdateStartUpdateCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsOtaUpdateStartUpdateCommand(buf []byte, offset flatbuffers.UOffsetT) *OtaUpdateStartUpdateCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OtaUpdateStartUpdateCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedOtaUpdateStartUpdateCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *OtaUpdateStartUpdateCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OtaUpdateStartUpdateCommand) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OtaUpdateStartUpdateCommand) Channel() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *OtaUpdateStartUpdateCommand) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func OtaUpdateStartUpdateCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func OtaUpdateStartUpdateCommandAddChannel(builder *flatbuffers.Builder, channel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(channel), 0)
}
func OtaUpdateStartUpdateCommandAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(version), 0)
}
func OtaUpdateStartUpdateCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
