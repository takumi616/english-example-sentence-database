package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"golang.org/x/sync/errgroup"
	"github.com/takumi616/go-english-vocabulary-api/config"
)

func run(ctx context.Context) error {
	//Handle signal
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()

	//Get config (environment variables)
	config, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return err
	}

	//Set up http listener with port from config file 
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("Failed to get a listener with port %d: %v", config.Port, err)
	}

	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Request Path: %s", r.URL.Path[1:])
		}),
	}

	//Run http server in another groutine
	//to stop it from external action
	//like sending cancel signal
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := server.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to run http server: %v", err)
			return err
		}
		return nil
	})

	//Wait for cancel notification 
	//ctx is canceled when groutine running http server returns error
	<-ctx.Done()
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Failed to shutdown http server: %v", err)
	}

	//Return goroutine's response (err or nil) 
	return eg.Wait()
}

func main() {
	err := run(context.Background())
	if err != nil {
		log.Printf("Main process error: %v", err)
	}
}