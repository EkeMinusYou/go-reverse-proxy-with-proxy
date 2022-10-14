package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	schema = "http"
	// backendHost = os.Getenv("BACKEND_HOST")
	proxyHost = os.Getenv("PROXY_HOST")
	port      = "8000"
)

func main() {
	url := &url.URL{
		Scheme: schema,
		Host:   proxyHost,
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = "example.com"
		req.Host = "example.com"
		dump, _ := httputil.DumpRequest(req, true)
		log.Println(string(dump))
	}
	handler := &httputil.ReverseProxy{
		Director:  director,
		Transport: tr,
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
