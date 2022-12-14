package heap

import (
	"jvmgo/ch11/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
	sourceFile        string
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceName()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fileds())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
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

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) JClass() *Object {
	return self.jClass
}
func (self *Class) getStaticMethod(name string, description string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == description {
			return method
		}
	}
	return nil
}

func (self *Class) Name() string {
	return self.name
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}
func (self *Class) getField(name string, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}
func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}
