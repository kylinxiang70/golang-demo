# http_tls 客户端校验

## 生成秘钥和证书

```
# 1. 生成CA的私钥
$ openssl genrsa -out ca.key 2048
Generating RSA private key, 2048 bit long modulus
.......................................+++
...............................+++
e is 65537 (0x10001)

# 2. 生成CA的数字证书
$ openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 365 -out ca.crt

# 3. 生成server的私钥
$ openssl genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus
.............................................................+++
..........................+++
e is 65537 (0x10001)

# 4. 生成server的数字证书请求server.csr
$ openssl req -new -key server.key -subj "/CN=localhost" -out server.csr

# 5. 生成server的数字证书
$ openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
Signature ok
subject=/CN=localhost
Getting CA Private Key
```

## 环境

golang 1.15之后 弃用了了x509 Common Name (name)
若需要开启，需要设置环境变量 GODEBUG=x509ignoreCN=0
> https://jishuin.proginn.com/p/763bfbd2a2ac

或者使用SAN证书
> https://www.cnblogs.com/jackluo/p/13841286.html
> http://liaoph.com/openssl-san/
## 代码
server.go
```go
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world\n")
	})

	err := http.ListenAndServeTLS(":8443", "./http_tls/certification/server.crt", "./http_tls/certification/server.key", nil)
	log.Fatal(err)
}
```

client.go
```go
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
```

## 运行

```bash
cd go-demo
go run ./http_tls/server/server.go
go run ./http_tls/client/client.go
```
