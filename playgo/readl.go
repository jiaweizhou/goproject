package playgo

import (
	"bytes"
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
