package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/server"
)

type Arith int

func (t *Arith) Mul(cxt context.Context, name string, reply *int) error {
	fmt.Println(123)
	return nil
}

func (t *Arith) Div(cxt context.Context, name string, reply *int) error {
	fmt.Println(123)

	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("ws", ":8972")
}
