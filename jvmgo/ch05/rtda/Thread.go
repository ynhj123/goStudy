package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
}
func (self *Thread) Pc() int {
	return self.pc
}
func (self *Thread) SetPc(pc int) {
	self.pc = pc
}
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals uint, maxStack uint) *Frame {
	return NewFrame(self, maxLocals, maxStack)
}
