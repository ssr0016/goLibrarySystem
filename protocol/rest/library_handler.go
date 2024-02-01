package rest

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/library/data/response"
	"github.com/ssr0016/library/helper"
	"github.com/ssr0016/library/library"
)

type Server struct {
	BookService library.Service
}

func NewServer(bookService library.Service) *Server {
	return &Server{
		BookService: bookService,
	}
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := s.BookService.Search(r.Context(), &library.SearchBookQuery{})
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cmd library.CreateBookCommand

	helper.ReadRequestBody(r, &cmd)

	s.BookService.Create(r.Context(), &cmd)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cmd,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (s *Server) GetById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId := ps.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result, err := s.BookService.GetById(r.Context(), int64(id))
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var cmd library.UpdateBookCommand
	helper.ReadRequestBody(r, &cmd)

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	helper.PanicIfError(err)
	cmd.ID = id

	s.BookService.Update(r.Context(), &cmd)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cmd,
	}

	helper.WriteResponseBody(w, webResponse)
}

// func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
// 	bookUpdateRequest := request.BookUpdateRequest{}
// 	helper.ReadRequestBody(requests, &bookUpdateRequest)

// 	bookId := params.ByName("bookId")
// 	id, err := strconv.Atoi(bookId)
// 	helper.PanicIfError(err)
// 	bookUpdateRequest.Id = id

// 	controller.BookService.Update(requests.Context(), bookUpdateRequest)
// 	webResponse := response.WebResponse{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	helper.WriteResponseBody(writer, webResponse)
// }

// func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
// 	bookId := params.ByName("bookId")
// 	id, err := strconv.Atoi(bookId)
// 	helper.PanicIfError(err)

// 	controller.BookService.Delete(requests.Context(), id)
// 	webResponse := response.WebResponse{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	helper.WriteResponseBody(writer, webResponse)
// }

// func (controller *BookController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	bookId := params.ByName("bookId")
// 	id, err := strconv.Atoi(bookId)
// 	helper.PanicIfError(err)

// 	result := controller.BookService.FindById(request.Context(), id)
// 	webResponse := response.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   result,
// 	}

// 	helper.WriteResponseBody(writer, webResponse)
// }

// func (controller *BookController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	result := controller.BookService.FindAll(request.Context())
// 	webResponse := response.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   result,
// 	}

// 	helper.WriteResponseBody(writer, webResponse)

// }
