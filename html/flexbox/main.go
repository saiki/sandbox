package main

import "log"
import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":80", nil))
}
