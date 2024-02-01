package main

import (
	"fmt"
	"net/http"

	"github.com/ssr0016/library/config"
	"github.com/ssr0016/library/helper"
	libraryimpl "github.com/ssr0016/library/library/libraryImpl"
	"github.com/ssr0016/library/protocol/rest"
	"github.com/ssr0016/library/router"
)

func main() {
	fmt.Printf("Start server!")

	//database
	db := config.DbConnectin()

	// repository
	bookStore := libraryimpl.NewStore(db)

	// service
	bookService := libraryimpl.NewService(bookStore)

	// controller
	bookHandler := rest.NewServer(bookService)
	// router
	routes := router.NewRouter(bookHandler)

	// routes := httprouter.New()

	// routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
