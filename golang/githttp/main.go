package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var root string

func main() {
	flag.StringVar(&root, "root", ".", "root dir.")
	flag.Parse()
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("root:", root)
	//	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	//		log.Println(req)
	//	})
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir(root))))
}
