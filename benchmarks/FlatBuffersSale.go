// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package dynobuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Sale struct {
	_tab flatbuffers.Table
}

func GetRootAsSale(buf []byte, offset flatbuffers.UOffsetT) *Sale {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Sale{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Sale) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Sale) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Sale) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Sale) Price() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Sale) MutatePrice(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *Sale) Quantity() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Sale) MutateQuantity(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func SaleStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SaleAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func SaleAddPrice(builder *flatbuffers.Builder, price float32) {
	builder.PrependFloat32Slot(1, price, 0.0)
}
func SaleAddQuantity(builder *flatbuffers.Builder, quantity int32) {
	builder.PrependInt32Slot(2, quantity, 0)
}
func SaleEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
