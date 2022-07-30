package main

import (
	"fmt"
	"jvmgo/ch04/classfile"
	"jvmgo/ch04/classpath"
	"strings"
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
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("cloud not find or load main class %s\n", cmd.class)
	//	return
	//}
	//fmt.Printf("class data:%v\n", classData)
	//fmt.Printf("classpath:%s calss:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceName())
	fmt.Printf("field count: %v\n", len(cf.Fileds()))
	for _, f := range cf.Fileds() {
		fmt.Printf("%s\n", f.Name())
	}
	fmt.Printf("method count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("%s\n", m.Name())
	}
}

func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

//func startJVM(options *cmdline.Option, class string,args []string) {
//
//}