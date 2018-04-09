package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello WORLD!!\n"))
       w.Write([]byte("AUTOMATION_TAG:NjxQnEw7EacICWoAXJrE31Dx6Vc8Kyx3JqtUqJu277tek8rcp3mJiblWZnjjqNbY"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
