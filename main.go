package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:PQ4MNS1z0dvIMEERqIpaQvYzXQH9.6Sj2vOvwRS93iCodosloia7wq3TnJ0heeBM"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
