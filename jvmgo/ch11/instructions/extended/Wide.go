package extended

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/instructions/loads"
	"jvmgo/ch11/instructions/math"
	"jvmgo/ch11/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		ints := &loads.ILOAD{}
		ints.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = ints
	case 0x16:
	case 0x17:
	case 0x18:
	case 0x19:
	case 0x36:
	case 0x37:
	case 0x38:
	case 0x39:
	case 0x3a:
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9:
		panic("Unsupported opcode:0xa9!")
	}
}
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
