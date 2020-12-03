/**
 * @author xiangqilin
 * @date 2020/12/1
**/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"unsafe"
)

func main() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	str := *(*string)(unsafe.Pointer(&body)) // convert byte[] to string
	fmt.Printf("%v", str)

}
