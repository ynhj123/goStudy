package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

func init() {
	native.Registry("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
