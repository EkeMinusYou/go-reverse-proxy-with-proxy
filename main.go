package main

import (
	"fmt"
	"net/http"
)

// var (
// 	schema    = "http"
// 	proxyHost = os.Getenv("PROXY_HOST")
// 	port      = "8000"
// )

// func main() {
// 	// handler := &httputil.ReverseProxy{}
// 	proxyUrl := &url.URL{
// 		Scheme: schema,
// 		Host:   proxyHost,
// 	}
// 	handler := httputil.NewSingleHostReverseProxy(proxyUrl)
// 	log.Fatal(http.ListenAndServe(
// 		":"+port,
// 		handler,
// 	))
// }

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
