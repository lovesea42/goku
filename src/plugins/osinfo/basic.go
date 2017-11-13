package osinfo

import (
	"runtime"
	"fmt"
)

type PluginTest struct{
	info map[string]string
}

func (self *PluginTest)Init(){
	fmt.Println("Plugin test init")
}

func (self *PluginTest)Collect() string{

	fmt.Println("Plugin test collect")

	return "test"
}

func OsTest(){
	GetOSArch()
}


/**
	获取操作系统架构
 */
func GetOSArch(){
	osarch := runtime.GOARCH
	fmt.Println(osarch)
}

func GetCurrentMemory(){
	//sysInfo := new(syscall.S)
}