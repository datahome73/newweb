package main 

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is my fisrt page! test deploy !!!</h1>")
	io.WriteString(w, "<h2>粤ICP备14074612</h2>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":80", nil)
}
