package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	if err := http.ListenAndServe("localhost:18080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
