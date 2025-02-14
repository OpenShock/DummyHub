// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Gateway

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Trigger struct {
	_tab flatbuffers.Table
}

func GetRootAsTrigger(buf []byte, offset flatbuffers.UOffsetT) *Trigger {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Trigger{}
	x.Init(buf, n+offset)
	return x
}

func FinishTriggerBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsTrigger(buf []byte, offset flatbuffers.UOffsetT) *Trigger {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Trigger{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedTriggerBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Trigger) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Trigger) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Trigger) Type() TriggerType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return TriggerType(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Trigger) MutateType(n TriggerType) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func TriggerStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TriggerAddType(builder *flatbuffers.Builder, type_ TriggerType) {
	builder.PrependByteSlot(0, byte(type_), 0)
}
func TriggerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
