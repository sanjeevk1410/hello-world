package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello WORLD!!\n"))
       w.Write([]byte("AUTOMATION_TAG:SoZgUV7WsjYolj36lmxkhk2MkF2w1mll2CnZhQXsQIBFtd4o2guMDSK11X2hfKqB"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
