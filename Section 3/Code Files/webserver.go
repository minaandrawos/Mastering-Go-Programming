package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080",nil)
}

func sroot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to the Hydra software system")
}
