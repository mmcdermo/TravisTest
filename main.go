package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/handlers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	z := "<HTML><HEAD><TITLE>INDEX</TITLE></HEAD><BODY>Hello, world v8!</BODY></HTML>"
	fmt.Fprintf(w, "%s", z)
}

func RunHTTPServer(port string) error{
	r := mux.NewRouter()
	r.HandleFunc("/index.html", IndexHandler)
	http.Handle("/", r)
	log.Println("Running HTTP Server on port " + port)
	err := http.ListenAndServe(":"+port, r)
	return err
}

func main(){
	err := RunHTTPServer("80")
	log.Println(err)
}
