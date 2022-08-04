package loads

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_FLOAD(frame, uint(self.Index))
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_FLOAD(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_FLOAD(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_FLOAD(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_FLOAD(frame, 3)
}

func _FLOAD(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}