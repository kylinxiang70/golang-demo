/**
 * @author xiangqilin
 * @date 2020/12/2
**/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("GODEBUG", "x509ignoreCN=0")
	fmt.Println(os.Getenv("GODEBUG"))
	caCrt, err := ioutil.ReadFile("./http_tls/certification/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	cerPool := x509.NewCertPool()
	cerPool.AppendCertsFromPEM(caCrt)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: cerPool,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))

}
