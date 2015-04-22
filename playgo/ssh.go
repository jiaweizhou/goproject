package playgo

import (
	"fmt"
	//"github.com/dynport/gossh"
	"golang.org/x/crypto/ssh"
	//"strconv"
	"bytes"
	"log"
	"strings"
	//"time"
)

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
func main() {
	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password("123456"))
	config := &ssh.ClientConfig{
		User: "root",
		Auth: auths,
	}
	conn, e := ssh.Dial("tcp", fmt.Sprintf("%s:%d", "10.10.102.249", 22), config)
	if e != nil {
		//log.Fatalf("unable to connect: %s", e)
		return
	}
	defer conn.Close()

	ses, e := conn.NewSession()
	if e != nil {
		return
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
	b := NewReadl()
	//in, e := ses.StdinPipe()
	//var out = make(chan bytes.Buffer, 10)
	//var b bytes.Buffer
	//var c []byte
	ses.Stdout = b
	in, e := ses.StdinPipe()
	//out, e := ses.StdoutPipe()
	go func() {
		for true {
			select {
			case line, _ := <-b.Out:
				fmt.Printf("%v", line.String())
				if strings.Contains(line.String(), "new UNIX password") {
					fmt.Fprintln(in, "123456")
				} else if strings.Contains(line.String(), "Is the information correct? [Y/n]") {
					fmt.Fprintln(in, "y")
				} else if strings.Contains(line.String(), "[]") {
					fmt.Fprintln(in, "")
				} else if strings.Contains(line.String(), "[Y/n]") {
					fmt.Fprintln(in, "y")
				} else if strings.Contains(line.String(), "[y/N]") {
					fmt.Fprintln(in, "y")
				} else if strings.Contains(line.String(), "already exists") {
				} else if strings.Contains(line.String(), "already a") {
				}
			}

			//fmt.Fscanln(out, b)
			//fmt.Printf("%v\n", b.String())
			//fmt.Println(b.String())
			//fmt.Fprintln(in, "123456")
			//b.String()
			//time.Sleep(100000000)
		}
	}()
	ses.Run("adduser zjw16;adduser zjw16 sudo")

	//fmt.Printf("%v\n", b.String())

}
