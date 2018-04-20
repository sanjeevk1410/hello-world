package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello WORLD!!\n"))
       w.Write([]byte("AUTOMATION_TAG:EzUe0RsEOAM.tgb8YS8O4swjkKJVzgkxpmpds9fgiPtfLM9t46QwDzvmeo1PS1Aa"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
