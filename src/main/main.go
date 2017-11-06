package main


/**

	GOOS=linux GOARCH=amd64 go build
 */
import (
	"config/tools"
	_"encoding/json"
	"plugins/jvm"
	"plugins"
	"config/options"
)

func testJvm(){
	data := new(jvm.ApiJvm)
	var i plugins.IPlugin = data
	i.Init()
	i.Collect()
}


func initOptions(){
	cfg := new(tools.Config)
	cfg.InitConfig("config.ini")
	defaultOptions := new(options.DefaultOptions)
	defaultOptions.Init(cfg)
}

func main(){

	initOptions()
	//testJvm()

}