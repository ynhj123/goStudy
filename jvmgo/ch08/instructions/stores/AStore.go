package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_ASTORE(frame, self.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_ASTORE(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_ASTORE(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_ASTORE(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_ASTORE(frame, 3)
}
func _ASTORE(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}
