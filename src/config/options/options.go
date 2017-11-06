package options

import (
	"config/tools"
)

const zoneconst string = "default"
const destServer_const string = "dest"
const interval_const string = "interval"


type DefaultOptions struct {
	destServer string		//上传服务器
	interval	int64			//采集间隔
}

func (self *DefaultOptions) Init(config *tools.Config){

	self.destServer = tools.ReadConfigString(config,zoneconst,destServer_const,"")
	self.interval = tools.ReadConfigInt64(config,zoneconst,interval_const,60000)

	println(self.destServer,self.interval)
}