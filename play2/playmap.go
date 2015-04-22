package main

import (
	//"encoding/json"
	"fmt"
	"reflect"
)

func a(re map[string]string) {
	re["aa"] = "asfsdfsd"
}
func main() {
	//var re map[string]string
	var re = map[string]string{}
	re["aa"] = "ggg"
	//data := `{"count":"abcd"}`
	rv := reflect.ValueOf(re)
	fmt.Println(rv.Kind())
	fmt.Println(rv.MapIndex(reflect.ValueOf("aa")))

	count := 1
	if value := reflect.ValueOf(count); value.CanSet() {
		value.SetInt(2)
	}
	fmt.Println(count)
	//err := json.Unmarshal([]byte(data), &re)
	//if err != nil {
	//panic(err)
	//}
}
