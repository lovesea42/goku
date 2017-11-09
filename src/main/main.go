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
	"flag"
	"constants"
	"upload"
)

func testJvm(url string){
	data := new(jvm.ApiJvm)
	var i plugins.IPlugin = data
	i.Init()
	collect := i.Collect()
	upload.HttpPost(url,collect)
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
		//正常命令执行
		//command.CreatePidFile(strconv.Itoa(os.Getpid()))

		options := initOptions()

		//ticker := time.NewTicker(time.Millisecond *
		//time.Duration(options.Interval))
		count:=1
		for true{
			println(count)
			testJvm(options.DestServer)
			time.Sleep(time.Millisecond *
				time.Duration(options.Interval))
			count++
		}


	}else{
		//特殊命令
		switch *s{

		case constants.CommandQuit:
			{
				//files := file.GetAllPid()
				//
				//for _,v := range files{
				//
				//}
			}
		default:


		}
	}



}