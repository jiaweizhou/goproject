// play2 project main.go
package main

import (
	"fmt"
	//"os"
	//"os/exec"
	//"deploy-engine/lib"
	//"path"
	//"strconv"
	//"strings"
	//"sync"
	//"encoding/json"
	//"gopkg.in/yaml.v2"
	//"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
func b(a sync.Mutex) {
	fmt.Println("asdgf")
}*/
func main() {
	var aa = `{ "people":[{ "firstName": "Brett", "lastName":"McLaughlin", "email": "aaaa" },{ "firstName": "Jason", "lastName":"Hunter", "email": "bbbb"},{ "firstName": "Elliotte", "lastName":"Harold", "email": "cccc" }]}`
	aa = `{"Components":{"nats":["10.10.10.1","10.10.10.2"],"haproxy":["10.10.10.3"]},
		"Properties":{"domain":"dawei.li"}
			}`
	/*
		c := bytes.Buffer{}
		c.WriteString(aa)
		fmt.Println("00")
		req, err := http.NewRequest("POST", "http://10.10.105.111:50000/deployment-engine/config", &c)
		if err != nil {
			fmt.Println("11")
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("22")
		}
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		//fmt.Println(string(res))
		defer res.Body.Close()
	*/
	var m = map[string]string{}
	fmt.Println(m["asdfsd"])
	t := strings.NewReader(aa)
	resp, _ := http.Post("http://10.10.101.184:50000/deployment-engine/status", "application/json", t)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//tt, _ := yaml.Marshal(t)
	//fmt.Println(t)
	/*a := sync.Mutex{}
	a.Lock()
	fmt.Println("asfasfsd")
	a.Unlock()
	fi, _ := os.Stat("/home/zjw/CF/cf170/cf_nise_installer/blobs/cache.gz")
	fmt.Println(fi.Size())
	//c := exec.Command("du -b #{filename} | awk '{print $1}'")
	//c.Start()
	srcpath := "/home/zjw/CF/cf170/cf_nise_installer/blobs/cache.gz"
	_, filename := path.Split(srcpath)
	filesize := fi.Size()
	//fmt.Println(filesize)
	s := lib.NewMyssh("10.10.101.181", "vcap", "password", "10.10.105.158", "zjw", "123456")
	fmt.Println(filename)
	fmt.Println(strings.Contains(s.Exec("ls"), filename))
	//strings.Contains()
	if strings.Contains(s.Exec("ls"), filename) {
		size, _ := strconv.Atoi(strings.TrimRight(s.Exec("du -b "+filename+" | awk '{print $1}'"), "\n\r"))
		if int64(size) == filesize {
			fmt.Println("YYYYYYYYYYYY")
		}
	}*/

}
