package playgo

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Job struct {
	Name string
	Temp map[string][]string
}
type Template struct {
	Deployment string
	Jobs       []Job
	Properties map[string]interface{}
}

type Config struct {
	Components map[string][]interface{}
	Properties map[string]string
}

func main() {
	t := Template{}
	inputFile := "template.yml"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	err = yaml.Unmarshal(buf, &t)
	w := Config{}
	inputFile = "config.yml"
	buf, err = ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	err = yaml.Unmarshal(buf, &w)

	switch t.Properties["domain"].(type) {
	case string:
		fmt.Println(t.Properties["domain"].(string))
	}
	fmt.Printf("%v", t.Properties["router"])
	//s := rein()

	switch t.Properties["router"].(type) {
	case map[interface{}]interface{}:
		pr := t.Properties["router"].(map[interface{}]interface{})
		p := pr["servers"]
		a := p.(map[interface{}]interface{})

		//var b []string
		//b :=
		a["z1"] = w.Components["haproxy"]

		switch a["z1"].([]interface{})[0].(type) {
		case string:
			fmt.Println("is string")
		}
		fmt.Printf("%v", a["z1"].([]interface{}))
		fmt.Println("Y")
	}
	as := map[string]string{}
	as["a"] = "a"

}

//ip = @configObj['components'][comp][0]
//newDomain = newDomain.sub(/<.+?>/, ip)
