package stores

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	_FSTORE(frame, self.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 3)
}
func _FSTORE(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
