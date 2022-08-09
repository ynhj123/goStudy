package misc

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	native.Registry("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jISysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jISysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
