package playgo

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	//"golang.org/x/crypto/ssh/agent"
	//"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strings"
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

type Scp struct {
	user      string
	password  string
	host      string
	localuser string
	localpwd  string
	localip   string
}

func (s *Scp) Upload(local string, dst string) {

	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(s.password))
	config := &ssh.ClientConfig{
		User: s.user,
		Auth: auths,
	}
	conn, e := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.host, 22), config)
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
		//log.Fatalf("request for pseudo terminal failed: %s", err)
	}
	readl := NewReadl()
	ses.Stdout = readl
	go func() {
		w, _ := ses.StdinPipe()
		defer w.Close()
		//content := "123456"
		for {
			select {
			case line, _ := <-readl.Out:
				fmt.Printf("%v", line.String())
				if strings.Contains(line.String(), "password") {
					fmt.Fprintln(w, s.localpwd)
				} else if strings.Contains(line.String(), "Are you sure you want to continue connecting") {
					fmt.Fprintln(w, "yes")
				}
				//fmt.Fprintln(w, "C0644", len(content), "testfile")
				//fmt.Fprintln(w, content)
				//fmt.Fprint(w, "\x00")
			}
		}

	}()
	if err := ses.Run("scp -r " + s.localuser + "@" + s.localip + ":" + local + " " + dst); err != nil {
		panic("Fail to run :" + err.Error())
	}
}

/*
func main() {
	scp := &Scp{
		localuser: "zjw",
		localpwd:  "123456",
		localip:   "10.10.105.158",
	}
	scp.upload("/home/zjw/adduser.rb", "~/")
	fi, err := os.Stat("../../goproject")
	if err != nil {
		fmt.Println("YY")
	}
	if fi.IsDir() {
		fmt.Println("NN")
	}
	dir, err := ioutil.ReadDir("../../goproject")
	if err != nil {
		panic("no dir:" + err.Error())
	}
	for _, fi := range dir {

		fmt.Println("%v", fi.Name())
	}

}
*/
