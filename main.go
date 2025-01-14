package main

import (
	"apiv2/cmd/api"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	// Capture interrupt signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	wg.Add(1)
	go api.StartServer(&wg, stopCh)

	// Wait for a termination signal
	<-signalCh
	fmt.Println("Shutting down server...")

	// Close the stop channel to notify the server
	close(stopCh)

	// Wait for the server goroutine to finish
	wg.Wait()
	fmt.Println("Server stopped.")
}
