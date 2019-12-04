package main

import (
	"fmt"
	"net/http"
)

func main(){
	// An http.ResponseWriter which is where you write your text/html response to.
	// An http.Request which contains all information about this HTTP request including things like the URL or header fields.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// An HTTP server has to listen on a port to pass connections on to the request handler
		// The following code will start Go’s default HTTP server and listen for connections on port 80.
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}