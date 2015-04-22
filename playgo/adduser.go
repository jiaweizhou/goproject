package playgo

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	Components map[string][]string
	Properties map[string]string
}
type Adduser struct {
	yml          *Config
	rootpassword string
	user         string
	userpassword string
}
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
func NewAdduser(configpath string) *Adduser {
	t := Config{}
	inputFile := "config.yml"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	err = yaml.Unmarshal(buf, &t)
	rootpassword := t.Properties["root"]
	user := t.Properties["user"]
	userpassword := t.Properties["password"]
	return &Adduser{
		yml:          &t,
		rootpassword: rootpassword,
		user:         user,
		userpassword: userpassword,
	}
}
func (t *Adduser) useradd(ip string, config *ssh.ClientConfig) {

	conn, e := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, 22), config)
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
					fmt.Fprintln(in, t.userpassword)
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
	ses.Run("adduser " + t.user + ";adduser " + t.user + " sudo")

	//fmt.Printf("%v\n", b.String())

}
func (t *Adduser) work() {
	isadduser, _ := strconv.Atoi(t.yml.Properties["adduser"])
	if isadduser == 0 {
		fmt.Println("Warn: config adduser 0. It means don't add.")
		return
	}
	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(t.rootpassword))
	config := &ssh.ClientConfig{
		User: "root",
		Auth: auths,
	}
	for _, ips := range t.yml.Components {
		for _, ip := range ips {
			t.useradd(ip, config)
		}
	}
}
func main() {
	t := NewAdduser("config.yml")
	if t.yml.Components["haproxy"] == nil {
		fmt.Println("Y")
	} else if t.yml.Components["haproxy"][0] == "" {
		fmt.Println("N")
	}
	ss := "<sun.test>"
	rep, _ := regexp.Compile(`<.+>`)
	c := rep.FindString(ss)
	fmt.Printf("%s\n", c)
	d := rep.FindSubmatch([]byte(ss))
	fmt.Printf("%s", string(d[0]))
	comp := rep.FindString(ss)
	comp = strings.TrimLeft(comp, "<")
	comp = strings.TrimRight(comp, ">")
	fmt.Printf("%s\n", comp)

}
