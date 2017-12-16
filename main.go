package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:6s8z5yZxpInKnnezG2W2V0itThbOfuXE9NmJyaQn92HFFgKe7bKras7WUF5CPXDl"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
