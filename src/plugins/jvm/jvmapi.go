package jvm

import (
	"os/exec"
	"bytes"
	"log"
	"regexp"
	"strconv"
	"fmt"
	"encoding/json"
)


type ApiJvm struct{
	info map[string]jvminfo		//具体信息集合
}

/**
	jvm详细信息
 */
type jvminfo struct{

	Eden 	float64		`json:"eden"`
	Old 	float64		`json:"old"`
	Ygc		int			`json:"ygc"`
	Ygct	float64		`json:"ygct"`
	Fgc		int			`json:"fgc"`
	Fgct 	float64		`json:"fgct"`

}

func (self *ApiJvm)Init(){
	self.info = make(map[string]jvminfo)
}

func (self *ApiJvm)Collect() string{

	jsonstr := make(map[string]string)

	//再修正
	apps := getAllJavaProcess()
	self.getJVMInfo(apps)

	for k,v := range self.info{
		b, _ :=  json.Marshal(v)
		jsonstr[k] = string(b)
	}

	b, _ := json.Marshal(jsonstr)
	return string(b)
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

	var slice []string = make([]string, 0)
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




/**
	获取jvm详细参数
 */
func (self *ApiJvm)getJVMInfo(apps []string){


	for i:= 0;i < len(apps);i++ {

		cmd := exec.Command("jstat","-gcutil",apps[i])
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil{
			log.Print(err)
			continue
		}

		out.ReadString('\n')
		gc,err := out.ReadString('\n')

		if err != nil{
			log.Print(err)
			continue;
		}

		pInfo := parserJvmInfo(gc)
		fmt.Println(pInfo)
		self.info[apps[i]] = *pInfo
	}

}

func parserJvmInfo(info string)*jvminfo{

	reg := regexp.MustCompile(`[0-9.-]+`)
	collect := reg.FindAllString(info,-1)

	//gcutil总长度为11
	if(len(collect) != 11){
		log.Print("error to parser jvm info,string length is" + strconv.Itoa(len(collect)) )
		return nil
	}

	ret := new(jvminfo)
	ret.Eden,_ = strconv.ParseFloat(collect[2], 32)
	ret.Old,_ = strconv.ParseFloat(collect[3], 32)
	ret.Fgc,_ = strconv.Atoi(collect[8])
	ret.Fgct,_= strconv.ParseFloat(collect[9], 32)
	ret.Ygc,_ = strconv.Atoi(collect[6])
	ret.Ygct,_= strconv.ParseFloat(collect[7], 32)

	return ret
}
