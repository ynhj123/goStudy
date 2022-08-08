package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"unsafe"
)

func init() {
	native.Registry("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Registry("java/lang/Object", "hashCode", "()I", hashCode)
	native.Registry("java/lang/Object", "clone", "()Ijava/lang/Object;", clone)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
