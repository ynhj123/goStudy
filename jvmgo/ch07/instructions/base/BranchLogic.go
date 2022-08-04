package base

import "jvmgo/ch07/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().Pc()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
