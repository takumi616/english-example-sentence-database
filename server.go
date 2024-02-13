package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv      *http.Server
	listener net.Listener
}

func NewServer(mux http.Handler, listener net.Listener) *Server {
	return &Server{
		srv:      &http.Server{Handler: mux},
		listener: listener,
	}
}

func (s *Server) Start(ctx context.Context) error {
	//Handle signal
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()

	//Run http server in another groutine
	//to stop it from external action
	//like sending cancel signal
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := s.srv.Serve(s.listener)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to run http server: %v", err)
			return err
		}
		return nil
	})

	//Wait for cancel notification
	//ctx is canceled when groutine running http server returns error
	<-ctx.Done()
	err := s.srv.Shutdown(context.Background())
	if err != nil {
		log.Printf("Failed to shutdown http server: %v", err)
	}

	//Return goroutine's response (err or nil)
	return eg.Wait()
}
