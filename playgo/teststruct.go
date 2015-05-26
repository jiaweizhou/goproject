package main

import (
	"fmt"
)

//var Appinfo = AppsInfo{}

type AppMetaInfo struct {
	Name     string
	Replicas int
	Status   int
}
type NamespaceInfo map[string]AppMetaInfo

//type AppsInfo map[string]NamespaceInfo

func main() {
	var app = AppMetaInfo{}
	var name = NamespaceInfo{}
	//app.Name = "asdfgsdfg"
	//Appinfo["a"] = map[string]*AppMetaInfo{}
	name["b"] = app
	//
	name["b"].Name = "asg"
	fmt.Println(name["b"].Name)
}
