# http tls 跳过客户端校验

## 使用openssl生成证书

`openssl genrsa -out server.key 2048 `

用于生成服务端私钥文件`server.key`，后面的参数2048单位是bit，是私钥的长度。

`openssl req -new -x509 -key server.key -out server.crt -days 365`

使用私钥生成服务端证书`server.crt`，有效期为365天

`server.key`和`server.crt`将作为`http.ListenAndServeTLS`的两个输入参数

## 代码

服务端代码client.go
```go
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
```
客户端代码
```go
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
```
