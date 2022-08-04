package main

import (
	"fmt"
	"jvmgo/ch07/instructions"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logInst)
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().Pc()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">>pc:%4d %v.%v%v\n", frame.NextPc(), className, method.Name(), method.Descriptor())
	}
}
