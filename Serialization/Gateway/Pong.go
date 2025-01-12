// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Gateway

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Pong struct {
	_tab flatbuffers.Table
}

func GetRootAsPong(buf []byte, offset flatbuffers.UOffsetT) *Pong {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Pong{}
	x.Init(buf, n+offset)
	return x
}

func FinishPongBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPong(buf []byte, offset flatbuffers.UOffsetT) *Pong {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Pong{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPongBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Pong) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Pong) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Pong) Uptime() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pong) MutateUptime(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Pong) Rssi() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pong) MutateRssi(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func PongStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PongAddUptime(builder *flatbuffers.Builder, uptime uint64) {
	builder.PrependUint64Slot(0, uptime, 0)
}
func PongAddRssi(builder *flatbuffers.Builder, rssi int32) {
	builder.PrependInt32Slot(1, rssi, 0)
}
func PongEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
