package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello World!</h1>"+"<div>Path: "+r.URL.Path+"</div>")
	log.Printf("Visit Path: %s\n", r.URL.Path)
}

func main() {

	// for {
	// 	fmt.Println("hello world")
	// 	time.Sleep(1 * time.Second)
	// }

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
