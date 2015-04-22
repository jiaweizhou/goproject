package main

import (
	//"bytes"
	"fmt"
	//"io"
	"os"
	"os/exec"
	//"time"
	"syscall"
)

type Test struct {
	s string
}

func PrintH() {
	fmt.Println("Hello")

}

var ch = make(chan int)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "scp doc.go root@10.10.101.184:~/")
	master, slave, err := os.Pipe()
	if err != nil {
		fmt.Println("pipe error:" + err.Error())
	}
	cmd.Stdin = master
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	err = cmd.Start()
	if err != nil {
		fmt.Println("start error:" + err.Error())
	}
	go func() {
		fmt.Println("out")
		//i := 0
		for {
			//i = i + 1
			//fmt.Println(i)
			//fmt.Println(i)
			slave.WriteString("123456\n")
			//time.Sleep(time.Second)
		}

	}()

	cmd.Wait()
	fmt.Println("outmain")
}
