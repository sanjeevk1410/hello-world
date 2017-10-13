package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:wXGPSrZiaHZ46HCX1QCi9F7DOhfAfUCWg8W3EzzQBXhQ2Fjl4v3lINQcwMc02iI1"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
