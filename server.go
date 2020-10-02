package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Responce struct {
	Time string `json:"time"`
}

func main() {
	fmt.Fprintf(w, "%s", "gogadoda")

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(Responce{time.Now().Format(time.RFC3339)})
		if err == nil {
			fmt.Fprintf(w, "%s", string(b))
		} else {
			fmt.Fprintf(w, "%s", err)
		}
	})
	http.ListenAndServe(":4444", nil)
}
