package playgo

import (
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
)

/*
func main() {
	c := exec.Command("scp", "/home/zjw/adduser.rb", "root@10.10.102.249:~/")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		//f.Write([]byte("foo\n"))
		//f.Write([]byte("bar\n"))
		//f.Write([]byte("baz\n"))
		//f.Write([]byte{4}) // EOT
		f.WriteString("123456")
	}()
	io.Copy(os.Stdout, f)
}
*/
