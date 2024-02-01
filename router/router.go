package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/library/protocol/rest"
)

func NewRouter(bookHandler *rest.Server) *httprouter.Router {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Welcome Home!")
	})

	router.GET("/api/book", bookHandler.Search)
	router.GET("/api/book/:bookId", bookHandler.GetById)
	router.POST("/api/book", bookHandler.Create)
	router.PATCH("/api/book/:bookId", bookHandler.Update)
	// router.DELETE("/api/book/:bookId", bookController.Delete)

	return router
}
