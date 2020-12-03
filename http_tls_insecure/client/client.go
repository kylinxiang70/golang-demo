/**
 * @author xiangqilin
 * @date 2020/12/2
**/
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			// 客户端默认开启对Server证书的校验，但Server证书不是由知名CA签发，这里跳过
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: transport,
	}

	resp, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
