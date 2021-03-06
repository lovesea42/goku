package command

import (
	"file"
	"os"
	"log"
	"constants"
)


func CheckPidExist(pid string)bool{
	return file.CheckFileIsExist(createPidName(pid))
}

func createPidName(pid string)string{
	return constants.PidFilePreConst + pid
}


func CreatePidFile(pid string){

	exist := file.CheckFileIsExist(pid)

	if !exist{
		_, err := os.Create(createPidName(pid))  //创建文件

		if err != nil{
			log.Panic("error to create pid file")
		}

	}


}
