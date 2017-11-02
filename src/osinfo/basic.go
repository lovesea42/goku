package osinfo

import (
	"runtime"
	"fmt"
)

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