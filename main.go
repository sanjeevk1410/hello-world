package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello WORLD!!\n"))
       w.Write([]byte("AUTOMATION_TAG:McC6capT8hWmmJ7Fvk1PkHrKbkYq21mmjTRZloF3K6jhG8xqR.NsmgrHacUcpwtf"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
