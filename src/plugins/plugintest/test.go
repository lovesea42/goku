package plugintest

import (
	"fmt"
)

func Test(){

}

type PluginTest struct{
	info map[string]string
}

func (self *PluginTest)Init(){
	fmt.Println("Plugin test init")
}

func (self *PluginTest)Collect() string{

	fmt.Println("Plugin test collect")

	return "test"
}