package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	//"io/ioutil"
	//"net/http"
	"strings"
)

func main() {
	//http.ServeFile(images.Ctx.ResponseWriter, images.Ctx.Request, imagepath)
	client := &http.Client{}
	body, _ := ioutil.ReadFile("request1.json")
	//var body = `{"name":"abcd","password":"123456"}`
	//fmt.Println(string(body))
	//reqest, _ := http.NewRequest("GET", "http://10.10.105.196:8080/v1/user/login", strings.NewReader(string(body)))
	reqest, _ := http.NewRequest("POST", "http://10.10.101.100:8080/api/v1beta3/namespaces/default/replicationcontrollers", strings.NewReader(string(body)))
	//reqest, _ := http.NewRequest("PUT", "http://10.10.103.86:8080/api/v1beta3/namespaces/default/services/logtest", strings.NewReader(string(body)))
	//reqest, _ := http.NewRequest("POST", "http://10.10.101.100:8080/api/v1beta3/namespaces/default/services", strings.NewReader(string(body)))
	//reqest, _ := http.NewRequest("PUT", "http://127.0.0.1:8080/v1/namespaces/np1/services/test/upgrade", strings.NewReader(string(body)))
	//reqest, _ := http.NewRequest("GET", "http://127.0.0.1:8080/v1/baseimages/8-jre8-customize.tar", strings.NewReader(""))
	//reqest, _ := http.NewRequest("PUT", "http://10.10.103.86:8080/api/v1beta3/namespaces/np1/replicationcontrollers/test-7", strings.NewReader(string(body)))
	//form := url.Values{"labelSelector": []string{"name=test1"}}
	//reqest.Form = form
	//reqest.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	//reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//reqest.Header.Set("Content-Type", "application/json")
	/*
		reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
		reqest.Header.Set("Cache-Control", "max-age=0")
	*/
	//reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	//response, _ := http.Get("http://121.40.171.96:8080/api/v1beta1/namespaces/default/replicationcontrollers?labelSelector=name%3Dfrontend")
	fmt.Println(response.StatusCode)
	body1, _ := ioutil.ReadAll(response.Body)
	//ioutil.WriteFile("8-jre8-customize.tar", body1, 0666)
	fmt.Println(string(body1))
	//}
}
