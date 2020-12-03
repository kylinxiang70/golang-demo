/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith struct{}

func (t *Arith) Multiply(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	if reply.C > 2<<31-1 || reply.C < -2<<31 {
		return errors.New("Exceeds the range of 32 - bit signed Numbers")
	}
	return nil
}
func main() {
	rpc.RegisterName("myArith", new(Arith))
	rpc.HandleHTTP()

	err := http.ListenAndServe(":18888", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
