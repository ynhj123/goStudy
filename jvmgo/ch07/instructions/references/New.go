package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPc()
		base.InitClass(frame.Thread(), class)
		return
	}
	if class.IsINTERFACE() || class.IsABSTRACT() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
