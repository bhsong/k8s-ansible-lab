package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello K8s! Response from Pod: %s\n", hostname)
	})

	fmt.Println("Server starting on port 8080...!")
	http.ListenAndServe(":8080", nil)
}