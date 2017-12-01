package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:uK.pZbHqmf4Or7e2iY67ODRUVzB.it.yclOrab5LdXLXkmYVoh2sm0WFk7Mezjor"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
