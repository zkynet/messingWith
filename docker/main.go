package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/echo", handlerEcho)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT:", r.RemoteAddr, " > ", r.RequestURI)
	fmt.Fprintln(w, "Hello docker users !")
}
func handlerEcho(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT:", r.RemoteAddr, " > ", r.RequestURI)
	fmt.Fprintln(w, "Hello echo !")
}
