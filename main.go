package main;

import (

	"net/http"
	"log"
	"io"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){

		//create the new request
		outreq := new(http.Request)
		*outreq = *req // includes shallow copies of maps, but okay
		outreq.URL.Scheme = "http"
		outreq.URL.Host = "localhost:8082" //new address

		//proxy transport
		transport := http.DefaultTransport
		transport.RoundTrip(outreq)

		//resend the request
		res, err := transport.RoundTrip(outreq)
		if err != nil {
			log.Print("http: proxy error: ", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		//copy the headers
		for k, vv := range res.Header {
			for _, v := range vv {
				resp.Header().Add(k, v)
			}
		}

		//copy the body
		io.Copy(resp, res.Body)
	});

	log.Print("started...")
	log.Fatal(http.ListenAndServe(":8081", nil));
}
