// addusertest.go
package test

import (
	"deploy-engine/lib"
	//"fmt"
)

func Addusertest() {
	t := lib.NewAdduser("./config.yml")
	t.Work()
	//fmt.Println("Hello World!")
}
