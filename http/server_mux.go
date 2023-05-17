package main

import "net/http"


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Tem mais controller do sistema
			w.WriteHeader(http.StatusOK)
		})

	mux.Handle("/blog", blog{})
	http.ListenAndServe(":8080", mux)

}

type blog struct {
	title string `json:"title"`
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))

}