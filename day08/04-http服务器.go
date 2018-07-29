package main

import (
	"net/http"
	"fmt"
)

func HandConn(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.RemoteAddr = ", r.RemoteAddr)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Body = ", r.Body)
	w.Write([]byte("hello world!"))
}

func main() {
	http.HandleFunc("/", HandConn)

	http.ListenAndServe(":8000", nil)
}
