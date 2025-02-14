// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Configuration

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CaptivePortalConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsCaptivePortalConfig(buf []byte, offset flatbuffers.UOffsetT) *CaptivePortalConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CaptivePortalConfig{}
	x.Init(buf, n+offset)
	return x
}

func FinishCaptivePortalConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsCaptivePortalConfig(buf []byte, offset flatbuffers.UOffsetT) *CaptivePortalConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CaptivePortalConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedCaptivePortalConfigBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *CaptivePortalConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CaptivePortalConfig) Table() flatbuffers.Table {
	return rcv._tab
}

/// Whether the captive portal is forced to be enabled
/// The captive portal will otherwise shut down when a gateway connection is established
func (rcv *CaptivePortalConfig) AlwaysEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Whether the captive portal is forced to be enabled
/// The captive portal will otherwise shut down when a gateway connection is established
func (rcv *CaptivePortalConfig) MutateAlwaysEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func CaptivePortalConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CaptivePortalConfigAddAlwaysEnabled(builder *flatbuffers.Builder, alwaysEnabled bool) {
	builder.PrependBoolSlot(0, alwaysEnabled, false)
}
func CaptivePortalConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
