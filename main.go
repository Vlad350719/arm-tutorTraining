package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"armwrestling/handlers/api"
)

func main() {
	http.HandleFunc("/api/user", api.UserHandler)
	http.HandleFunc("/api/react", api.ReactHandler)

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	log.Println("start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
