/**
 * @author xiangqilin
 * @date 2020/12/3
**/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	caCrt, err := ioutil.ReadFile("./http_mtls/certification/ca.crt")
	if err != nil {
		log.Fatal(caCrt)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("./http_mtls/certification/client.crt",
		"./http_mtls/certification/client.key")
	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caPool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
	}

	resp, err := c.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(body)
	}
	fmt.Println(string(body))
}
