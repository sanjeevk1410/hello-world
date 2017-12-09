package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:L2nuUV8SdAKZPxyKq709.iptS6PpJd8t3mwACUR6XT6DA2L3UWxy9okSOHJqamBG"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
