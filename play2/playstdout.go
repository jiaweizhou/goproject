package main

import (
	"fmt"
	"os"
	"syscall"
	//"io/ioutil"
	//"bufio"
	"bytes"
	//"io"
	//"strings"
	"errors"
)

type PipeReader interface {
	Read(p []byte) (n int, err error)
}
type PipeWriter interface {
	Write(p []byte) (n int, err error)
}

var ch = make(chan int)

func PipeCopy(dst PipeWriter, src PipeReader) (written int64, err error) {
	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = errors.New("short write")
				break
			}
		}
		if nr < 32*1024 {
			break
		}
		if er == errors.New("EOF") {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
func main() {
	//stdout := os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	t := &bytes.Buffer{}

	r, w, err := os.Pipe()
	if err != nil {
		panic(err.Error())
	}

	go func(w *os.File) {
		//syscall.Close(zz[0])

		//syscall.Dup2(syscall.Stdout, int(w.Fd()))
		os.Stdout = w
		fmt.Println("NNNN")
		fmt.Println("asdfasdg")
		ch <- 1
	}(w)
	<-ch
	os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	fmt.Println("abc")
	var wait = make(chan int, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				w.WriteString("abcdefgdddddddddddddddddd\n")
			}
			wait <- 1
		}()
	}
	for i := 0; i < 10; i++ {
		<-wait
	}
	PipeCopy(t, r)
	fmt.Println(t.String())

	//w.WriteString("abc")
	//w.WriteString("abc")
	//r.Read(t)
	//fmt.Println("%v", t)
	//os.Stdout.Read(t)
	//fmt.Println(string(t))
	//syscall.Close(file)
	//syscall.Read(syscall.Stdout, t)
	//fmt.Println("%v", string(t))
	//syscall.Read(syscall.Stderr, t)
	//fmt.Println("%v", string(t))

}
