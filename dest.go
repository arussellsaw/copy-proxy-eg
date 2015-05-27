package main
import (
	"net/http"
	"log"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		log.Print(req);
		resp.Header().Add("X-I-Like-Cats", "true")
		resp.Write([]byte("DEST REACHED"));

	})

	log.Fatal(http.ListenAndServe(":8082", nil));
}
