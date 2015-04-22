package playgo

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"time"
	"path"
)

type Config struct {
	Components map[string][]string
	Properties map[string]string
}
type FileSync struct {
	localPath  string
	remotePath string
	cfg 	&Config
	remoteIp string
	password string
	localTime time.Timer
	dirs []string
}
func NewFlieSync(ip string ,remotepath string) *FileSync{
	t := Config{}
	inputFile := "config.yml"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	err = yaml.Unmarshal(buf, &t)
	return &FileSync{
		localPath:	"../..",
		remotePath:	remotepath,
		cfg:			&t,
		remoteIp:	ip==""? t.Properties["home"] : ip,
		password:	t.Properties["password"],
	}
}
