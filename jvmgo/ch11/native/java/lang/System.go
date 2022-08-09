package lang

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

func init() {
	native.Registry("java/lang/System", "arraycopy", "(Ljava/lang/Object;Ljava/lang/Object;II)V;", arraycopy)
}

func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	if src == nil || dest == nil {
		panic("java.lang.NullPointException")
	}
	if checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || destPos < 0 || length < 0 || srcPos+length > dest.ArrayLength() || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutBoundsException")
	}
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}
func setOut(frame *rtda.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/printStream;", out)

}
