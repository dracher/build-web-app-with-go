package main

import (
	// "encoding/xml"
	"fmt"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

type Profile struct {
	Name   string
	Hobbie []string
}

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome. visit sub page now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./templates called "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HTML(w, http.StatusOK, "example", nil)
	})
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", mux)
}
