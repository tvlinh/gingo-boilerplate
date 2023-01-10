package handlerhttp

import (
	"tvlinh/gingo-boilerplate/internal/middleware"
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/handler/http"
	"tvlinh/gingo-boilerplate/internal/modules/helloworld/usecase"

	"github.com/gin-gonic/gin"
)

func HandlerHelloWorld(r *gin.RouterGroup, mw *middleware.MiddlewareManager) {
	handlerHello("", r, mw)
}

func handlerHello(path string, r *gin.RouterGroup, mw *middleware.MiddlewareManager) {
	helloUC := usecase.NewHello()
	helloHandle := http.NewHello(helloUC)

	r.GET(path, helloHandle.Hello())
}
