package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s\n", data)
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.ListenAndServe(":9090", nil)
}
