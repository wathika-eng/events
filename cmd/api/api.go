package api

import (
	"apiv2/pkg/db"
	"apiv2/pkg/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const PORT string = ":8080"

func StartServer(wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()

	// Initialize the server
	server := gin.Default()
	db.InitDB()
	routes.RegisterRoutes(server)

	// Create an HTTP server
	httpServer := &http.Server{
		Addr:    PORT,
		Handler: server,
	}

	// Goroutine to start the server
	go func() {
		fmt.Printf("Server listening on http://localhost%v\n", PORT)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", PORT, err)
		}
	}()

	// Wait for the stop signal
	<-stopCh

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server stopped gracefully.")
}
