package http

import (
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/presenter"
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/usecase"

	"github.com/gin-gonic/gin"
)

type HelloAdvance interface {
	Hello() gin.HandlerFunc
}

type helloAdvance struct {
	uc usecase.HelloAdvance
}

func (h *helloAdvance) Hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hello := h.uc.Hello()

		result := &presenter.Hello{}
		ctx.JSON(200, result.FromModel(hello))
	}
}

func NewHelloAdvance(uc usecase.HelloAdvance) HelloAdvance {
	return &helloAdvance{uc: uc}
}
