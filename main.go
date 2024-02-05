package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	server := &http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Request Path: %s", r.URL.Path[1:])
		}),
	}

	//Run http server in another groutine
	//to stop it from external action
	//like sending cancel signal
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to run http server: %v", err)
			return err
		}
		return nil
	})

	//Wait for cancel notification 
	//ctx is canceled when groutine running http server returns error
	<-ctx.Done()
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Failed to shutdown http server: %v", err)
	}

	//Return goroutine's response (err or nil) 
	return eg.Wait()
}

func main() {
	//Create run() to make it easy testing http server
	err := run(context.Background())
	if err != nil {
		log.Printf("Main process error: %v", err)
	} else {
		log.Println("Main process finished successfully.")
	}
}