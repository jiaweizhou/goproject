package main

import (
	//"bufio"
	//"bytes"
	"fmt"
	//"io"
	//"os"
	//"os/exec"
	"syscall"
)

var ch = make(chan int)

func main() {
	/*
		//fmt.Println(os.Environ())
		in := &bytes.Buffer{}
		in.WriteString("123456\n")

		cmd := exec.Command("/usr/bin/scp", "sshtool", "root@10.10.101.175:~/")
		stdin, _ := cmd.StdinPipe()
		stdout, _ := cmd.StdoutPipe()
		//stderr, _ := cmd.StderrPipe()
		outbr := bufio.NewReader(stdout)

		//errbr := bufio.NewReader(stderr)
		go func() {

			os.Stdin.WriteString("123456\n")
			stdin.Write([]byte("123456"))
			t, _, _ := outbr.ReadLine()
			fmt.Println(string(t))
			fmt.Println("YYYYY")
			ch <- 1
		}()
		cmd.Start()
		<-ch
	*/
	err := syscall.Termios

	if err != nil {
		fmt.Println(err.Error())
	}
}
