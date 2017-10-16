package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:OTukfDNBDHB8jXny3QyfQ.Ig6A1q.Qsv0AA4sWSIyzh949uXS8.JPzsQtjqXAJ8p"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
