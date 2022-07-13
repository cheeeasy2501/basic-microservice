package v1

import (
	"basic-microservice/internal/entity"
	"basic-microservice/internal/service"
	"basic-microservice/internal/transport/http/form"
	"basic-microservice/pkg/httpserver"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newBookRoutes(r *gin.RouterGroup, s service.IBookService) {
	h := newBookHandler(s)

	g := r.Group("/books")
	{
		g.GET("/", h.GetBooks)
		g.GET("/:id", h.GetBook)
		g.POST("/", h.CreateBook)
		g.PATCH("/", h.UpdateBook)
		g.DELETE("/", h.DeleteBook)
	}
}

type BookHandler struct {
	s service.IBookService
}

func newBookHandler(s service.IBookService) *BookHandler {
	return &BookHandler{
		s: s,
	}
}

func (r *BookHandler) GetBooks(ctx *gin.Context) {

}

func (r *BookHandler) GetBook(ctx *gin.Context) {

}

// return entity.Book ???
func (r *BookHandler) CreateBook(ctx *gin.Context) {
	f := &form.CreateBookForm{}
	err := ctx.ShouldBindJSON(f)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpserver.NewErrorResponse("Binding error", err)) /// method to gin response
		return
	}

	formErrorResponse := f.LoadAndValidate()
	if formErrorResponse != nil {
		ctx.JSON(http.StatusBadRequest, formErrorResponse)
		return
	}

	// todo: work with bookService and return entity or httpserver.Response
	book := entity.BookEntity{
		Title:       f.Title,
		Description: f.Description,
	} // mock
	book, err = r.s.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpserver.NewErrorResponse("Book isn't created", err))
	}

	ctx.JSON(http.StatusOK, httpserver.NewResponse("Book created", book))
	return
}

func (r *BookHandler) UpdateBook(ctx *gin.Context) {

}

func (r *BookHandler) DeleteBook(ctx *gin.Context) {

}
