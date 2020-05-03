package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	address = "localhost"
	port    = "8080"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run() {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error while starting server: %v", err)
	}
}

func (s *Server) Shutdown() {
	log.Print("gracefully shutting down API...")

	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error while shutting down server: %v", err)
	}
}

func CreateServer() *Server {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	return &Server{srv: &http.Server{
		Addr:    fmt.Sprintf("%s:%s", address, port),
		Handler: router,
	}}
}

func CreateSignalChannel() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	return done
}

func main() {
	myServer := CreateServer()

	go myServer.Run()

	done := CreateSignalChannel()
	<-done

	myServer.Shutdown()
}
