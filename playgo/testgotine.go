package playgo

import (
	"fmt"
)

var c = make(chan int)

func abc() {

	a := 2
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("ttttt")
			a = 1
			c <- 1
		}()
	}
	for i := 0; i < 10; i++ {
		<-c
	}
	fmt.Println(a)
}
func main() {

	abc()

}
