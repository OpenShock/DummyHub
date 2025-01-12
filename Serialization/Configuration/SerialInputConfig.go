// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Configuration

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SerialInputConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsSerialInputConfig(buf []byte, offset flatbuffers.UOffsetT) *SerialInputConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SerialInputConfig{}
	x.Init(buf, n+offset)
	return x
}

func FinishSerialInputConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSerialInputConfig(buf []byte, offset flatbuffers.UOffsetT) *SerialInputConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SerialInputConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSerialInputConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SerialInputConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SerialInputConfig) Table() flatbuffers.Table {
	return rcv._tab
}

/// Whether to echo typed characters back to the serial console
func (rcv *SerialInputConfig) EchoEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// Whether to echo typed characters back to the serial console
func (rcv *SerialInputConfig) MutateEchoEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func SerialInputConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SerialInputConfigAddEchoEnabled(builder *flatbuffers.Builder, echoEnabled bool) {
	builder.PrependBoolSlot(0, echoEnabled, true)
}
func SerialInputConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
