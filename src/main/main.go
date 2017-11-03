package main


/**

	GOOS=linux GOARCH=amd64 go build
 */
import (
	//"jvm"
	//"osinfo"
	//"file"
	"config"
	"fmt"
	"encoding/json"
)

func main(){

	cfg := new(config.Config)
	cfg.InitConfig("config.ini")
	fmt.Println(cfg.Read("default","dest"))
	//osinfo.GetOSArch()
	//jvm.Test()
	//file.FileCreator()
}