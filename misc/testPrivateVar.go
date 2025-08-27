package demo

import "fmt"

type private struct {
	S int
}

func NewPrivate() *private {
	return &private{}
}

func (p *private) TestPricvate() {
	p.S = 1
	fmt.Printf("TestPricvate p: %+v", p)
}
