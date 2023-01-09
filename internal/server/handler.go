package server

import (
	"tvlinh/gingo-boilerplate/internal/middleware"
	"tvlinh/gingo-boilerplate/internal/server/handlerhttp"
)

func (g *GinServer) MapHandlers() error {
	v1 := g.engine.Group("/api/" + g.cfg.Server.Version)
	mw := middleware.NewMiddlewareManager()

	g.engine.Use(mw.Base(), mw.Cors())

	handlerhttp.HandlerHelloWorld(v1.Group("/helloworld"), mw)
	return nil
}
