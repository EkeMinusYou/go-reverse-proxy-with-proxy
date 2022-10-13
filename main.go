package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	schema    = "http"
	proxyHost = os.Getenv("PROXY_HOST")
	port      = 8000
)

func main() {
	// handler := &httputil.ReverseProxy{}
	proxyUrl := &url.URL{
		Scheme: schema,
		Host:   proxyHost,
	}
	handler := httputil.NewSingleHostReverseProxy(proxyUrl)
	log.Fatal(http.ListenAndServe(
		":"+string(rune(port)),
		handler,
	))
}
