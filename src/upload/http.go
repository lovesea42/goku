package upload

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
)

func HttpPost(url string,content string){

	var jsonstr =[]byte(content)

	req,err := http.NewRequest("POST",url,bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type","application/json")

	client := &http.Client{}
	resp,err := client.Do(req)

	if err != nil{
		println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Println(string(body))
}
