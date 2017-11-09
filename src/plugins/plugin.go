package plugins

import (
)

/**
	插件接口
 */

type IPlugin interface{
	Init()
	Collect() string
}


//未完成，插件模式加载
func LoadPlugins(){

}