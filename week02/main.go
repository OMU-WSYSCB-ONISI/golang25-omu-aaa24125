package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hellohandler)
	http.ListenAndServe(":8080", nil)
}

/* 以下，関数を追加 */
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}
