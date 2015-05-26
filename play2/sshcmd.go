package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	//"golang.org/x/crypto/ssh/terminal"
	"io"
	//"io/ioutil"
	"bufio"
	"log"
	//"os"
	//"strings"
	//"bufio"
)

var ch = make(chan int)

type Readl struct {
	Out chan bytes.Buffer
}

func (r *Readl) Write(p []byte) (n int, err error) {
	t := bytes.Buffer{}
	t.Write(p)
	//fmt.Printf("%v", t.String())
	r.Out <- t
	return len(p), nil
}
func NewReadl() *Readl {
	t := make(chan bytes.Buffer, 1000)
	return &Readl{
		Out: t,
	}
}
func cmd(host string, user string, password string, cmd string) {
	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(password))
	config := &ssh.ClientConfig{
		User: user,
		Auth: auths,
	}
	//connect
	conn, e := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, 22), config)
	if e != nil {
		panic("Fail to connect :" + e.Error())
	}
	//c.Result[key+host] = append(c.Result[key+host], key+" ["+host+"]:")
	defer conn.Close()
	//exec
	ses, e := conn.NewSession()
	defer ses.Close()
	if e != nil {
		panic("Fail to create session :" + e.Error())
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal

	if err := ses.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}
	stdout, err := ses.StdoutPipe()
	//stdout := NewReadl()
	//ses.Stdout = stdout

	stdin, err := ses.StdinPipe()
	if err != nil {
		fmt.Println("unable to acquire stdin pipe: %s", err)
	}
	err = ses.Shell()
	fmt.Println("aaaaa")
	if err != nil {
		fmt.Println("session failed: %s", err)
	}
	go func() {
		stdin.Write([]byte(cmd + " && exit \n"))
		//stdin.Write([]byte("exit\n"))
		//stdin.Write([]byte("exit\n"))
		fmt.Println("bbbb")
		//var buf bytes.Buffer
		/*
			if _, err := io.Copy(&buf, stdout); err != nil {
				fmt.Println(err.Error())
			}*/
		//fmt.Println("bbbb")
		//fmt.Println("%s", buf.String())
		out := bufio.NewReader(stdout)
		for {
			line, _, err := out.ReadLine()
			if err == io.EOF {
				break
			}
			fmt.Println(string(line))
		}
		ch <- 1
	}()
	<-ch
}
func main() {
	cmd("10.10.101.184", "root", "123456", "top -n 1")
}
