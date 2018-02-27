package LearnGo

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"io/ioutil"
)

func Ioutil(){
	link := "www.baidu.com"
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport:tr}
	response,err := client.Get(link)
	if err != nil{
		fmt.Println(err)
	}
	defer response.Body.Close()
	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}