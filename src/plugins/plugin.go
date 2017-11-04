package plugins

import (
)

/**
	插件接口
 */

type IPlugin interface{
	Init()
	Collect()
}


//未完成，插件模式加载
func LoadPlugins(){

}