package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!!\n"))
       w.Write([]byte("AUTOMATION_TAG:Uf9uBvdIhubxjs.oxi9g0Rtw841MWukTP8t2SJ9t12FggHtNuN8Ovusl8d0HcA.F"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
