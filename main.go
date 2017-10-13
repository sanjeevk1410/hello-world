package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!!\n"))
        w.Write([]byte("AUTOMATION_TAG:LysUogn0Ba2Zy.YlwArC7IZTnHUJH9nBA98gdyeD7j.GpR5sQkGGwN5ws5MP47xM"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
