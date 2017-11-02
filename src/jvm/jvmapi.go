package jvm

import (
	"os/exec"
	"bytes"
	"log"
	"regexp"
)


func Test(){
	slice := getAllJavaProcess()
	slice.
	println(slice)
}

/**
	获取所有JAVA进程
 */
func getAllJavaProcess()[]string{

	cmd := exec.Command("jps")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		log.Panic(err)
	}

	var slice []string = make([]string, 10)
	for{
		line,err := out.ReadString('\n')

		if err != nil{
			break;
		}

		reg := regexp.MustCompile(`[0-9]+`)
		reg2 := regexp.MustCompile(`[a-zA-Z]+`)
		name := reg2.FindString(line)
		if(name == "Jps"){
			continue;
		}
		slice = append(slice,reg.FindString(line))
	}

	return slice
}