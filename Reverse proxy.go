package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type handle struct {
	host string
	port string
}

type Service struct {
	auth *handle
	user *handle
}

func (this *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var remote *url.URL
	if strings.Contains(r.RequestURI, "api/auth") {
		remote, _ = url.Parse("http://" + this.auth.host + ":" + this.auth.port)
	} else if strings.Contains(r.RequestURI, "api/user") {
		remote, _ = url.Parse("http://" + this.user.host + ":" + this.user.port)
	} else {
		fmt.Fprintf(w, "404 Not Found")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	// 注册被代理的服务器 (host， port)
	service := &Service{
		auth: &handle{host: "127.0.0.1", port: "8081"},
		user: &handle{host: "127.0.0.1", port: "8082"},
	}
	err := http.ListenAndServe(":8888", service)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
