package main

import (
	"fmt"
	//"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var uname string
	fmt.Fprintf(w, "Hi there, Welcome to GoChatter!")
	// io.WriteString(w, "This is my website!\n")
	fmt.Fprintf(w, "Enter your username: ")
	fmt.Scan(&uname)
	fmt.Println(uname)
}

func main() {
	http.HandleFunc("/chat", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
