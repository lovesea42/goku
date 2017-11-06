package options

import (
	"config/tools"
)

const zoneconst string = "default"
const destServer_const string = "dest"
const interval_const string = "interval"


type DefaultOptions struct {
	DestServer string		//上传服务器
	Interval	int64			//采集间隔
}

func (self *DefaultOptions) Init(config *tools.Config){

	self.DestServer = tools.ReadConfigString(config,zoneconst,destServer_const,"")
	self.Interval = tools.ReadConfigInt64(config,zoneconst,interval_const,60000)

	//println(self.DestServer,self.Interval)
}