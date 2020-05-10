package http

import (
	"fmt"
	"net/http"
)

// IndexHandler api
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// SampleCommunication http handle
func SampleCommunication(url string, port string) {

	// httpHandler.SampleCommunication("/", ":3001")
	http.HandleFunc(url, IndexHandler)
	http.ListenAndServe(port, nil)

}
