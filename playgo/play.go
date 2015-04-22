/*
package main

import (
	"fmt"
)

func modify(a string) {
	a = "aa"
}
func inff(b interface{}) {
	switch b.(type) {
	case int:
		fmt.Printf("%d", b.(int))
	}
}
func main() {

	var t map[string][]string
	inff(t)
	a := "adsfg"
	modify(a)
	fmt.Printf("%s", a)
}
*/
package main

import (
	"fmt"
	"os"
)

func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		fmt.Printf("%v", fi.Mode())
		return fi.IsDir()
	}

	panic("not reached")
}

func main() {
	if isDirExists("./aa") {
		fmt.Println("目录存在")
	} else {
		fmt.Println("目录不存在")
	}
	//os.MkdirAll("test2", 0777)

}
