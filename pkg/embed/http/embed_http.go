package main

import (
	"embed"
	"net/http"
)

var (
	//go:embed hello.txt
	f embed.FS
)

func main() {
	http.Handle("/", http.FileServer(http.FS(f)))
	http.ListenAndServe(":8080", nil)
}
