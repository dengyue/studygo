package main

import "net/http"

func f1(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/hello", f1)
	http.ListenAndServe("127.0.0.1:8080", nil)
}