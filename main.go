package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:dwKzA3NUv0Vv3EMpRYjg0bCbQ2JHl9nvtPDBOMxsiL3T6jlo4W.nxaL.2Ex82wPG"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
