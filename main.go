package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	proxyScheme   = os.Getenv("PROXY_SCHEME")
	proxyHost     = os.Getenv("PROXY_HOST")
	backendSchema = os.Getenv("BACKEND_SCHEME")
	backendHost   = os.Getenv("BACKEND_HOST")
	port          = "8000"
)

func main() {
	url := &url.URL{
		Scheme: proxyScheme,
		Host:   proxyHost,
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	director := func(req *http.Request) {
		req.URL.Scheme = backendSchema
		req.URL.Host = backendHost
		req.Host = backendHost
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
