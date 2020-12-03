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
	client, err := rpc.DialHTTP("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err.Error())
	}
	args := Args{2 << 35, 2}
	reply := Reply{}

	err = client.Call("myArith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(reply)
}
