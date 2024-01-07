package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")

		if err != nil { 
			fmt.Println(err)
		}

		errorFormatted := fmt.Sprintf("Number of Bytes written: %d", n)
		fmt.Println(errorFormatted)
	})
	
	http.ListenAndServe(":8080", nil)
}