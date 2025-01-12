// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Local

import (
	flatbuffers "github.com/google/flatbuffers/go"

	OpenShock__Serialization__Types "OpenShock/Serialization/Types"
)

type WifiNetworkEvent struct {
	_tab flatbuffers.Table
}

func GetRootAsWifiNetworkEvent(buf []byte, offset flatbuffers.UOffsetT) *WifiNetworkEvent {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WifiNetworkEvent{}
	x.Init(buf, n+offset)
	return x
}

func FinishWifiNetworkEventBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsWifiNetworkEvent(buf []byte, offset flatbuffers.UOffsetT) *WifiNetworkEvent {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WifiNetworkEvent{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedWifiNetworkEventBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *WifiNetworkEvent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WifiNetworkEvent) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WifiNetworkEvent) EventType() OpenShock__Serialization__Types.WifiNetworkEventType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return OpenShock__Serialization__Types.WifiNetworkEventType(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *WifiNetworkEvent) MutateEventType(n OpenShock__Serialization__Types.WifiNetworkEventType) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *WifiNetworkEvent) Networks(obj *OpenShock__Serialization__Types.WifiNetwork, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *WifiNetworkEvent) NetworksLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func WifiNetworkEventStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WifiNetworkEventAddEventType(builder *flatbuffers.Builder, eventType OpenShock__Serialization__Types.WifiNetworkEventType) {
	builder.PrependByteSlot(0, byte(eventType), 0)
}
func WifiNetworkEventAddNetworks(builder *flatbuffers.Builder, networks flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(networks), 0)
}
func WifiNetworkEventStartNetworksVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func WifiNetworkEventEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
