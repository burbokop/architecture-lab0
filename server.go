package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Responce struct {
	Time string `json:"time"`
}

func main() {
	m := Responce{"Alice"}
	b, err := json.Marshal(m)
	fmt.Println(string(b), err)

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		m := Responce{"Alice"}
		b, err := json.Marshal(m)
		fmt.Println(string(b))
		fmt.Fprintf(w, "Hello World! time %s %s %s", b, string(b), err)
	})
	http.ListenAndServe(":4444", nil)
}
