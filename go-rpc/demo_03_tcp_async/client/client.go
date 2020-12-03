/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {
	args := Args{5, 6}
	var reply Reply

	client, err := rpc.Dial("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("result: %v.", reply)
}
