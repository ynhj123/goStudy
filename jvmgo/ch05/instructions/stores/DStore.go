package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	_DSTORE(frame, self.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_DSTORE(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_DSTORE(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_DSTORE(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_DSTORE(frame, 3)
}
func _DSTORE(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
