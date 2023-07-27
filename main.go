package main

import (
	"goweb/myapp"
	"net/http"
)



func main() {
	// http
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello bar")
	// })

	// http.HandleFunc("/zoo", zooHandler)
	// http.Handle("/foo", &footHandler{})
	// http.ListenAndServe(":3000",nil )

	// mux
	
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}