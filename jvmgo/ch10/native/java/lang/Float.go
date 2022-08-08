package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
	"math"
)

func init() {
	native.Registry("java/lang/Float", "floatToRawInBits", "(F)I", floatToRawInBits)
}

func floatToRawInBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
