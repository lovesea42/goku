package file

import (
	"os"
	"io/ioutil"
	"fmt"
	"regexp"
	"constants"
)

func CheckFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func GetAllPid()[]string{

	ret := make([]string, 0)

	dir_list,e := ioutil.ReadDir("/")

	if e != nil {
		fmt.Println("read regex file path error")
		return ret
	}

	r1,_ := regexp.Compile(constants.PidFilePreConst + "[0-9]+")
	r2,_ := regexp.Compile("[0-9]+")

	for _,v := range dir_list {


		match := r1.FindString(v.Name())

		if(len(match) > 0){
			pid := r2.FindString(match)
			ret = append(ret, pid)
		}
	}

	return ret
}