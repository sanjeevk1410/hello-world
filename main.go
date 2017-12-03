package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:UN3abJNHZchladjlI5IZBaoJO5gp7b.u6R8h4Nt6ecw0eoOSVjix1AeH6W7uiKfo"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
