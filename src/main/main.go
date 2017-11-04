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
	_"encoding/json"
	"plugins/jvm"
	"plugins"
)

func testJvm(){
	data := new(jvm.ApiJvm)
	var i plugins.IPlugin = data
	i.Init()
	i.Collect()
	//slice := getAllJavaProcess()
	//data.getJVMInfo(slice)
}

func main(){

	cfg := new(config.Config)
	cfg.InitConfig("config.ini")
	fmt.Println(cfg.Read("default","dest"))
	//osinfo.GetOSArch()
	testJvm()
	//file.FileCreator()
}