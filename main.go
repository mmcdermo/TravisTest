package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/handlers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	z := "<HTML><HEAD><TITLE>INDEX</TITLE></HEAD><BODY>I should be on dev only!! Update 245667778901!</BODY></HTML>"
	fmt.Fprintf(w, "%s", z)
}

func RootHandler(w http.ResponseWriter, r *http.Request){
	z := "<HTML><HEAD><TITLE>INDEX</TITLE></HEAD><BODY>!</BODY></HTML>"
	fmt.Fprintf(w, "%s", z)
}

func RunHTTPServer(port string) error{
	r := mux.NewRouter()
	r.HandleFunc("/index.html", IndexHandler)
	r.HandleFunc("/index2.html", IndexHandler)
	r.HandleFunc("/", IndexHandler)
	http.Handle("/", r)
	log.Println("Running HTTP Server on port " + port)
	err := http.ListenAndServe(":"+port, r)
	return err
}

func main(){
	err := RunHTTPServer("80")
	log.Println(err)
}
