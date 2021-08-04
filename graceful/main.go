package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const addr = ":8082"

	r := gin.Default()

	// TODO: Setup middlewares: CORS, Logs, etc
	// TODO: Register routes

	log.Println("start listen and serve at " + addr)
	s := &http.Server{Addr: addr, Handler: r}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve service: %v\n", err)
			return
		}
	}()

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
		return
	}
}
