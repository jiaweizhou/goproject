package playgo

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"

	"os"
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

type Config struct {
	Components map[string][]string
	Properties map[string]string
}
type Command struct {
	yml    *Config
	thread []chan int
	cmd    string
	result map[string][]string
	job    string
	ip     string
}

func NewCommand(path string, cmd string, job string, ip string) *Command {
	t := Config{}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	err = yaml.Unmarshal(buf, &t)
	return &Command{
		yml:    &t,
		cmd:    cmd,
		job:    job,
		ip:     ip,
		result: map[string][]string{},
		thread: []chan int{},
	}
}
func (c *Command) exec(key string, host string, user string, password string) {
	fmt.Printf("%s%s\n", user, password)
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
	c.result[key+host] = []string{}
	c.result[key+host] = append(c.result[key+host], key+" ["+host+"]:")
	defer conn.Close()

	//exec
	ses, e := conn.NewSession()
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
	readl := NewReadl()
	ses.Stdout = readl
	go func() {
		w, _ := ses.StdinPipe()
		defer w.Close()
		for {
			select {
			case line, _ := <-readl.Out:
				//fmt.Printf("%v", line.String())
				if strings.Contains(line.String(), "[sudo]") {
					fmt.Fprintln(w, password)
				} else {
					c.result[key+host] = append(c.result[key+host], strings.Trim(line.String(), "\n"))
				}
				//fmt.Fprintln(w, "C0644", len(content), "testfile")
				//fmt.Fprintln(w, content)
				//fmt.Fprint(w, "\x00")
			}
		}

	}()
	fmt.Println(host)
	if err := ses.Run(c.cmd); err != nil {
		panic("Fail to run :" + err.Error())
	}
	fmt.Println(host)
}
func (c *Command) get_through() {
	for key, hosts := range c.yml.Components {
		for _, host := range hosts {
			user := c.yml.Properties["user"]
			password := c.yml.Properties["password"]
			//fmt.Println(key, host, user, password)
			ch := make(chan int)
			c.thread = append(c.thread, ch)
			go func(key string, host string, user string, password string, ch chan int) {
				//fmt.Println("sadasd")
				defer func() { // 必须要先声明defer，否则不能捕获到panic异常
					//fmt.Println("c")
					if err := recover(); err != nil {
						fmt.Printf("Job: %s on %s exec %s ERROR.\n", key, host, c.cmd)
						fmt.Println(err)
						ch <- 1
					}
					//ch <- 1
					//fmt.Println("d")
				}()
				//fmt.Println("sadasd")
				if c.job != "" && key != c.job {
					ch <- 0
					return
				}
				if c.ip != "" && host != c.ip {
					ch <- 0
					return
				}
				//fmt.Println("%s", host)
				//c.thread = append(c.thread, ch)
				c.exec(key, host, user, password)
				ch <- 1

			}(key, host, user, password, ch)
		}
	}
	fmt.Println("%d", len(c.thread))

	for _, v := range c.thread {
		<-v
	}
}
func (c *Command) show() {
	for _, v := range c.result {
		fmt.Println(v)
	}
}
func main() {
	cmd := NewCommand("config.yml", "ls", "", "")
	//fmt.Println("%v", cmd.yml.Components)

	//cmd.result["aa"] = []string{}
	//a := map[string]string{}
	//a["asdf"] = "asg"
	cmd.get_through()
	cmd.show()
}
