/**
 * @author xiangqilin
 * @date 2020/12/2
**/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("GODEBUG", "x509ignoreCN=0")
	fmt.Println(os.Getenv("GODEBUG"))
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world\n")
	})

	err := http.ListenAndServeTLS(":8443", "./http_tls/certification/server.crt", "./http_tls/certification/server.key", nil)
	log.Fatal(err)
}
