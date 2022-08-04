package rtda

import "jvmgo/ch07/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	method       *heap.Method
	thread       *Thread
	nextPc       int
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPc() int {
	return self.nextPc
}

func (self *Frame) SetNextPC(nextPc int) {
	self.nextPc = nextPc
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) RevertNextPc() {
	self.nextPc = self.thread.pc
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
