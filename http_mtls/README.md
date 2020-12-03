# http mtls

## 生成私钥和证书

```bash
$ openssl genrsa -out ca.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
.....+++++
...+++++
e is 65537 (0x010001)

$ openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 365 -out

$ openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 365 -out ca.crt

$ openssl genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
............................+++++
..+++++
e is 65537 (0x010001)

$ openssl req -new -key server.key -subj "/CN=localhost" -out server.csr

$ openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
Signature ok
subject=CN = localhost
Getting CA Private Key

$ openssl genrsa -out client.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
................................................+++++
........................+++++
e is 65537 (0x010001)

$ openssl req -new -key client.key -subj "/CN=localhost" -out client.csr

$ openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365
Signature ok
subject=CN = localhost
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
```