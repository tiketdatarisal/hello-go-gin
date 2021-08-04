package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tiketdatarisal/hello-go-gin/recipes/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const addr = ":8082"

	// Initialize Gin
	r := gin.Default()

	// Configure and setup CORS middleware
	cfg := cors.DefaultConfig()
	cfg.AllowWildcard = true
	cfg.AllowOrigins = []string{"http://*.tiket.com", "https://*.tiket.com"}
	cfg.AddAllowHeaders("Authorization", "Content-Type")
	r.Use(cors.New(cfg))

	// Register routes
	routesCtx := handlers.Context{Gin: r}
	routesCtx.RegisterRoutes()

	// Start receiving request
	log.Println("start listen and serve at " + addr)
	s := &http.Server{Addr: addr, Handler: r}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve service: %v\n", err)
			return
		}
	}()

	// Wait until someone/something terminate this service
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	// Shutting down server
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
		return
	}
}
