package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! time")
	})
	http.ListenAndServe(":4444", nil)
}
