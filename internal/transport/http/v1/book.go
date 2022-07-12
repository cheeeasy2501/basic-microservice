package v1

import (
	"basic-microservice/internal/transport/http/form"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newBookRoutes(r *gin.RouterGroup) {
	h := newBookHandler()

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
}

func newBookHandler() *BookHandler {
	return &BookHandler{}
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
		ctx.JSON(http.StatusInternalServerError, err) /// method to gin response
		return
	}

	errResp := f.LoadAndValidate()
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	// todo: work with bookService and return entity or httpserver.Response
	return
}

func (r *BookHandler) UpdateBook(ctx *gin.Context) {

}

func (r *BookHandler) DeleteBook(ctx *gin.Context) {

}
