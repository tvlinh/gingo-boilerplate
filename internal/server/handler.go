package server

import (
	"tvlinh/gingo-boilerplate/internal/middleware"
	"tvlinh/gingo-boilerplate/internal/server/handlerhttp"
)

func (s *GinServer) MapHandlers() error {
	v1 := s.engine.Group("/api/" + s.cfg.Server.Version)
	mw := middleware.NewMiddlewareManager()

	s.engine.Use(mw.Base(), mw.Cors())

	handlerhttp.HandlerHelloWorld(v1.Group("/helloworld"), mw)
	return nil
}
