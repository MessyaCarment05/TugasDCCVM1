package main

import (
	"fmt"
	"net/http"
)

func VM1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ini adalah Virtual Machine 1")
}
func main(){
	http.HandleFunc("/", VM1)
	http.ListenAndServe(":8000", nil)
}