package main

import (
	"net/http"
)

func main() {

    http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
    http.HandleFunc("/sang", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello sang"))
    })
    http.HandleFunc("/slime", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello slime"))
    })
    http.HandleFunc("/yhong", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello yhong"))
    })
  	http.HandleFunc("/Byounghak", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello Byounghak"))
	  })
    http.HandleFunc("/jung", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello jung"))
    })
    http.HandleFunc("/Jerry", func(w http.ResponseWriter, req *http.Request) {
      w.Write([]byte("Hello Jerry"))
    })
 
    http.ListenAndServe(":5000", nil)

}
