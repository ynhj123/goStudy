package main

import (
	"fmt"
	"jvmgo/ch04/rtda"
)

func main() {
	var cmd = parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testOperandStack(stack *rtda.OperandStack) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(2997924580)
	stack.PushLong(-2997924580)
	stack.PushFloat(3.1415926)
	stack.PushDouble(2.71828182845)
	stack.PushRef(nil)
	println(stack.PopRef())
	println(stack.PopDouble())
	println(stack.PopFloat())
	println(stack.PopLong())
	println(stack.PopLong())
	println(stack.PopInt())
	println(stack.PopInt())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}
