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
	"time"
	"main/command"
	"os"
	"strconv"
	"flag"
)

func testJvm(){
	data := new(jvm.ApiJvm)
	var i plugins.IPlugin = data
	i.Init()
	i.Collect()
}


func initOptions()*options.DefaultOptions{
	cfg := new(tools.Config)

	cfg.InitConfig("config.ini")

	defaultOptions := new(options.DefaultOptions)
	defaultOptions.Init(cfg)

	return defaultOptions
}

func main(){

	s := flag.String("s", "", "command")
	flag.Parse()

	if *s != ""{

		command.CreatePidFile(strconv.Itoa(os.Getpid()))

		options := initOptions()

		//ticker := time.NewTicker(time.Millisecond *
		//time.Duration(options.Interval))
		count:=1
		for true{
			println(count)
			testJvm()
			time.Sleep(time.Millisecond *
				time.Duration(options.Interval))
			count++
		}


	}else{

	}



}