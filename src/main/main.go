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
	"plugins/redis"
	"fmt"
)

func testJvm(url string){
	data := new(jvm.ApiJvm)
	var i plugins.Pluginer = data
	i.Init()
	collect := i.Collect()
	upload.HttpPost(url,collect)
}

func testRedis(url string){

	conn,err := redis.Connect(url)
	if(err != nil){
		fmt.Println("error to connect redis " + url)
		return
	}

	reply,err := conn.Do("info")

	if(err != nil){
		fmt.Println("error to do ping " + err.Error())
		return
	}

	fmt.Println(reply)

}


func initOptions()*options.DefaultOptions{
	cfg := new(tools.Config)

	cfg.InitConfig("config.ini")

	defaultOptions := new(options.DefaultOptions)
	defaultOptions.Init(cfg)

	return defaultOptions
}

//暂时无法实现
func loadPlugins(){


	//pkg,err := importer.Default().Import("osinfo")
	//if err != nil {
	//	fmt.Printf("error: %s\n", err.Error())
	//	return
	//}
	//
	//t1 := reflect.ValueOf("daw").Type()
	//t2 := reflect.New(t1).Elem()
	//fmt.Println(t2)
	//
	//t := reflect.ValueOf("PluginTest").Type()
	//v := reflect.New(t).Elem()
	//test := v.Interface().(plugins.Pluginer)
	////reflect.ValueOf(v).MethodByName("Init").Call(nil)
	//test.Init()
	//
	//for _, declName := range pkg.Scope().Names() {
	//	//reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	//	fmt.Println(declName)
	//	}
	//
	//	fmt.Println("end to load Plugins")
}

func main(){

	//loadPlugins()

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
			//testJvm(options.DestServer)
			testRedis("192.168.100.144:9000")
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