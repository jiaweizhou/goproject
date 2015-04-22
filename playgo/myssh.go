package playgo

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"log"
	//"os"
	"strconv"
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
	t := make(chan bytes.Buffer, 10)
	return &Readl{
		Out: t,
	}
}

/*
type Scp struct {
	user      string
	password  string
	host      string
	localuser string
	localpwd  string
	localip   string
}*/
type Myssh struct {
	user     string
	password string
	ip       string
	client   *ssh.Client
	Scpcli   *Scp
}

func NewMyssh(ip string, user string, password string, localip string, localuser string, localpwd string) *Myssh {
	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(password))
	config := &ssh.ClientConfig{
		User: user,
		Auth: auths,
	}
	conn, e := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, 22), config)
	if e != nil {
		//log.Fatalf("unable to connect: %s", e)
		panic(e.Error())
	}
	return &Myssh{
		client:   conn,
		user:     user,
		password: password,
		ip:       ip,
		Scpcli: &Scp{
			user:      user,
			password:  password,
			host:      ip,
			localuser: localuser,
			localpwd:  localpwd,
			localip:   localip,
		},
	}
}
func (m *Myssh) exec(cmd string) string {
	var res string

	ses, e := m.client.NewSession()
	if e != nil {
		return ""
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
	readl := NewReadl()
	ses.Stdout = readl
	go func() {
		w, _ := ses.StdinPipe()
		defer w.Close()
		for {
			select {
			case line, _ := <-readl.Out:
				//fmt.Printf("%v", line.String())
				res = res + line.String()
				if strings.Contains(line.String(), "[sudo]") {
					fmt.Fprintln(w, m.password)
				} else {
					//c.result[key+host] = append(c.result[key+host], strings.Trim(line.String(), "\n"))
				}
				//fmt.Fprintln(w, "C0644", len(content), "testfile")
				//fmt.Fprintln(w, content)
				//fmt.Fprint(w, "\x00")
			}
		}

	}()
	//fmt.Println(host)
	if err := ses.Run(cmd); err != nil {
		panic("Fail to run :" + err.Error())
	}
	return res
	//fmt.Println(host)
}

/*
func main() {
	mys := NewMyssh("10.10.102.249", "root", "123456", "10.10.105.158", "zjw", "123456")
	z := mys.exec("du -b adduser.rb | awk '{print $1}'")
	fmt.Println("%v", []byte(z))
	fmt.Printf("%v", string(13) == "\n")

	z = strings.TrimRight(z, "\n\r")

	fmt.Printf("%v", []byte(z))
	t, err := strconv.Atoi(z)
	if err != nil {
		fmt.Println("tttt")
	}
	fmt.Printf("%d", t)
}
*/
