package main

import (
	"fmt"

	"github.com/DavidHuie/quartz/go/quartz"
)

type Adder struct{}

type Args struct {
	A int
	B int
}

type Response struct {
	Sum int
}

func (t *Adder) Add(args Args, response *Response) error {
	*response = Response{args.A + args.B}
	return nil
}

func main() {
	fmt.Println("Hello from Golang!")
	myAdder := &Adder{}
	quartz.RegisterName("my_adder", myAdder)

	quartz.Start()
}
