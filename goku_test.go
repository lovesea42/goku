package goku

import (
	"testing"
	"config"
	"fmt"
)

func TestGoku(t *testing.T){
	cfg := new(config.Config)
	cfg.InitConfig("config.ini")
	fmt.Println(cfg.Read("default","dest"))

	//osinfo.GetOSArch()
	//testJvm()
}
