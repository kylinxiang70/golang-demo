/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package main

import (
	"errors"
	"log"
	"net"
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
	rpc.RegisterName("Arith", new(Arith))
	listener, err := net.Listen("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		go func(conn net.Conn) {
			rpc.ServeConn(conn)
		}(conn)
	}
}
