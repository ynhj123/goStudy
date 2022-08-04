package comparisons

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_DCMP(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_DCMP(frame, false)
}

func _DCMP(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
