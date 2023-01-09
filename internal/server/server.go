package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tvlinh/gingo-boilerplate/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinServer struct {
	engine *gin.Engine
	cfg    *config.Config
	ready  chan bool
	logger *logrus.Logger
}

func NewGinServer(cfg *config.Config, logger *logrus.Logger, ready chan bool) *GinServer {
	return &GinServer{
		engine: gin.Default(),
		cfg:    cfg,
		ready:  ready,
		logger: logger,
	}
}

func (s *GinServer) Run() error {
	if err := s.MapHandlers(); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":" + s.cfg.Server.Port,
		Handler: s.engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("listen: %s\n", err)
		}
	}()

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	defer s.logger.Fatalln("Server Exited Properly With Err ", ctx.Err())

	return srv.Shutdown(ctx)
}
