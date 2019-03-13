package main

import (
	"fmt"
	"net/http"
)

func getstatus(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Status: Okay")
}

func main(){
	http.HandleFunc("/status", getstatus)
	http.ListenAndServe(":8080", nil)
}


