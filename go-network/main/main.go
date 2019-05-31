package main

import (
	"net/http"
	"fmt"
)

func headler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World , %s!", request.URL.Path[1:])
}


func main(){
    http.HandleFunc("/", headler)
    http.ListenAndServe(":8080", nil)
}