package tools

import (
	"os"
	"log"
	"bufio"
	"strings"
	"io"
	"errors"
)

/**
	配置模块
 */
type IOption interface{
	Init()
}



const middle = "========="

//配置信息
type Config struct {
	configMap map[string]string
	strcet string
}

func (c *Config)InitConfig(path string){
	c.configMap = make(map[string]string)

	f,err := os.Open(path)
	if err != nil{
		log.Panic("error to read confi" + path)
	}

	    defer f.Close()

	r:=bufio.NewReader(f)

	for{
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.configMap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) (string,error) {
	key = node + middle + key
	v, found := c.configMap[key]
	if !found {
		return "",errors.New("error to read file")
	}
	return v,nil
}