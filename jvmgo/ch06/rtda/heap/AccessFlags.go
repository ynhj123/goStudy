package heap

import "jvmgo/ch06/classfile"

const (
	ACC_PUBLIC       = 0x0001
	ACC_PRIVATE      = 0x0002
	ACC_PROTECTED    = 0x0004
	ACC_STATIC       = 0x0008
	ACC_FINAL        = 0x0010
	ACC_SUPER        = 0x0020
	ACC_SYNCHRONIZED = 0x0020
	ACC_VOLATILE     = 0x0040
	ACC_BRIDGE       = 0x0040
	ACC_TRANSIENT    = 0x0080
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_INTERFACE    = 0x0200
	ACC_ABSTRACT     = 0x0400
	ACC_STRICT       = 0x0800
	ACC_SYNTHETIC    = 0x1000
	ACC_ANNOTATION   = 0x2000
	ACC_ENUM         = 0x4000
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceName()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fileds())
	class.methods = newMethods(class, cf.Methods())
	return class
}
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsPRIVATE() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Class) IsPROTECTED() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *Class) IsSTATIC() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Class) IsFINAL() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSUPER() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsSYNCHRONIZED() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}

func (self *Class) IsVOLATILE() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}

func (self *Class) IsBRIDGE() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}

func (self *Class) IsTRANSIENT() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}

func (self *Class) IsVARARGS() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Class) IsNATIVE() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Class) IsINTERFACE() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsABSTRACT() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSTRICT() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Class) IsSYNTHETIC() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsANNOTATION() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsENUM() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
