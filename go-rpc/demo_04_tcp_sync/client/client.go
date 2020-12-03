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
	call := client.Go("Arith.Multiply", &args, &reply, nil)
	fmt.Println("async test text...")
	<-call.Done
	fmt.Printf("result: %v.", reply)
}
