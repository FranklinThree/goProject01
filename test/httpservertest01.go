package main

import (
	"fmt"
	"net/http"
)
func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello world")
}
func main(){
	http.HandleFunc("/",HelloHandler)
	http.ListenAndServe("127.0.0.1:33333",nil)
}
