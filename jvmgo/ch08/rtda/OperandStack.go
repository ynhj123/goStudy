package rtda

import (
	"jvmgo/ch08/rtda/heap"
	"math"
)

type OperandStack struct {
	size uint
	Slot []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{Slot: make([]Slot, maxStack)}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32) {
	self.Slot[self.size].num = val
	self.size++
}
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.Slot[self.size].num
}
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.Slot[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.Slot[self.size].num)
	return math.Float32frombits(bits)
}
func (self *OperandStack) PushLong(val int64) {
	self.Slot[self.size].num = int32(val)
	self.Slot[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.Slot[self.size].num)
	hign := uint32(self.Slot[self.size+1].num)
	return int64(hign)<<32 | int64(low)
}
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.Slot[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	return self.Slot[self.size].ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.Slot[self.size] = slot
	self.size++
}
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.Slot[self.size]

}

func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.Slot[self.size-1-n].ref
}