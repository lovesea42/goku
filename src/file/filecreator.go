package file

import (
	"os"
	"fmt"
	"flag"
	"math/rand"
	"time"
	"strconv"
)

func writeRandomFile(name string,size int){
	userFile := name //文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	//循环写入数据到文件
	for i:=0;i<size;i++{
		var slice []byte = make([]byte, 1024)
		fout.Write(slice)//强转成byte slice后再写入
	}
}


func writeFile(min int,max int,size int){
	for i:=0;i<size;i++{
		sz := rand.Intn(max - min) + min
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		filename := tm.Format("20060102030405") + strconv.Itoa(rand.Intn(100000000))
		writeRandomFile(filename,sz)
	}
}

func FileCreator(){

	fmt.Println("---------本程序为文件生成工具-------------")
	fmt.Println("命令介绍：")
	fmt.Println("-l 最小文件大小,如-l=100 最小生成100K大小文件")
	fmt.Println("-m 最大文件大小,如-m=400 最小生成400K大小文件")
	fmt.Println("-s 文件数量，如-s 10000，生成10000个文件")
	fmt.Println("--------------------------------------------")

	min := flag.Int("l", 100, "min file size")
	max := flag.Int("m", 400, "max file size")
	size := flag.Int("s", 1000, "file count")

	flag.Parse()

	writeFile(*min,*max,*size)
	//osinfo.GetOSArch()
	//jvm.Test()

}