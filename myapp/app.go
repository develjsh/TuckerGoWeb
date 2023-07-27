package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type footHandler struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (f *footHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func zooHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	
	fmt.Fprintf(w, "Hello %s", name)
}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type userHandler struct{}

func (u  *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	 }

	user.CreatedAt = time.Now()

	data, err := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("bar")
		fmt.Fprint(w, "Hello bar")
	})

	http.HandleFunc("/zoo", zooHandler)
	mux.Handle("/foo", &footHandler{})
	mux.Handle("/user", &userHandler{})

	return mux
}