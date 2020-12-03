/**
 * @author xiangqilin
 * @date 2020/12/1
**/
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "hello world\n")
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServeTLS(":8443", "./http_tls_insecure/certification/server.crt", "./http_tls_insecure/certification/server.key", nil)
	if err != nil {
		panic(err)
	}
}
