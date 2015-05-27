package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Host = "localhost:8082"
			},
		}
		proxy.ServeHTTP(resp, req)
	})

	log.Print("started...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
