package tools

import "strconv"

/**

	读取配置信息

 */
func ReadConfigString(config *Config,zonename string,key string,dft string) string {

	cfg := *config

	ret,err := cfg.Read(zonename,key)
	if err != nil{
		ret = dft
	}

	return ret
}

func ReadConfigInt64(config *Config,zonename string,key string,dft int64) int64 {

	var rt int64;
	cfg := *config
	ret,err1 := cfg.Read(zonename,key)

	if err1 == nil {
		ps,err2 := strconv.ParseInt(ret,10,64)
		if err2 != nil{
			rt = dft
		}else{
			rt = ps
		}

	}else{
		rt = dft
	}

	return rt
}