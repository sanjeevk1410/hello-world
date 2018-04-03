package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Hello WORLD!!\n"))
       w.Write([]byte("AUTOMATION_TAG:KXPm.WXKkZEQGa49zfv7BZA7h9BAIdLLWNI9T14T6G.Ng1CQzl9qzRSV4chk7CUH"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
