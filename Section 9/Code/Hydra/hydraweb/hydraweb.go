package hydraweb

import (
	"Hydra/hlogger"
	"fmt"
	"net/http"
	//"time"
	"log"
)

func Run() {
	http.HandleFunc("/", sroot)
	http.Handle("/testhandle", newHandler())
	http.HandleFunc("/testquery", queryTestHandler)
	http.ListenAndServe(":8080", nil)
	/*
		server := &http.Server{
			Addr:         ":8080",
			Handler:      newHandler(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
		server.ListenAndServe()
	*/
}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Forms", r.Form)
	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v \n", q)

	v1, v2 := q.Get("key1"), q.Get("key2")
	if v1 == v2 {
		message = message + fmt.Sprintf("V1 and V2 are equal %s \n", v1)
	} else {
		message = message + fmt.Sprintf("V1 is equal %s, V2 is equal %s \n", v1, v2)
	}
	fmt.Fprint(w, message)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra software system")
	logger.Println("Received an http Get request on root url")
}
