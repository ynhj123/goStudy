package io

import (
	"jvmgo/ch11/rtda"
	"os"
	"unsafe"
)

func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	jBytes := b.Data().([]int8)
	goBytes := castInt8ToUint8(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}

func castInt8ToUint8(jBytes []int8) (goBytes []byte) {
	prt := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(prt))
	return
}
