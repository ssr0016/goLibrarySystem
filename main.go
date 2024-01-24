package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/library/helper"
)

func main() {
	fmt.Printf("Start server!")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Hello, World!")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
