// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Gateway

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ShockerCommandList struct {
	_tab flatbuffers.Table
}

func GetRootAsShockerCommandList(buf []byte, offset flatbuffers.UOffsetT) *ShockerCommandList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ShockerCommandList{}
	x.Init(buf, n+offset)
	return x
}

func FinishShockerCommandListBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsShockerCommandList(buf []byte, offset flatbuffers.UOffsetT) *ShockerCommandList {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ShockerCommandList{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedShockerCommandListBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ShockerCommandList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ShockerCommandList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ShockerCommandList) Commands(obj *ShockerCommand, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ShockerCommandList) CommandsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ShockerCommandListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ShockerCommandListAddCommands(builder *flatbuffers.Builder, commands flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(commands), 0)
}
func ShockerCommandListStartCommandsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ShockerCommandListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
