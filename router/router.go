package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/library/protocol/rest"
)

func NewRouter(bookController *rest.BookController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Welcome Home!")
	})

	router.GET("/api/book", bookController.FindAll)
	router.GET("/api/book/:bookId", bookController.FindById)
	router.POST("/api/book", bookController.Create)
	router.PATCH("/api/book/:bookId", bookController.Update)
	router.DELETE("/api/book/:bookId", bookController.Delete)

	return router
}
