package main

import (
	"fmt"
	"jvmgo/ch11/classpath"
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func (self JVM) start() {
	self.initVM()
	self.execMain()
}

func (self JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

func (self JVM) execMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class%s\n", self.cmd.class)
		return
	}
	argArr := self.createArgArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argArr)
	self.mainThread.PushFrame(frame)

	interpret(mainMethod, self.cmd.verboseInstFlag)

}

func (self JVM) createArgArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}
func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}
