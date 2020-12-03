/**
 * @author xiangqilin
 * @date 2020/12/3
**/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	caCrt, err := ioutil.ReadFile("./http_mtls/certification/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCrt)
	s := http.Server{
		Addr:    ":8443",
		Handler: HelloHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  caPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("./http_mtls/certification/server.crt",
		"http_mtls/certification/server.key")
	if err != nil {
		log.Fatal(err)
	}
}
