package http

import (
	"github.com/gin-gonic/gin"
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/presenter"
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/usecase"
)

type Hello interface {
	Hello() gin.HandlerFunc
}

type hello struct {
	uc usecase.Hello
}

func (h *hello) Hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hello := h.uc.Hello()

		result := &presenter.Hello{}
		ctx.JSON(200, result.FromModel(hello))
	}
}

func NewHello(uc usecase.Hello) Hello {
	return &hello{uc: uc}
}
