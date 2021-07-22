package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	bye_sigs := make(chan bool, 1)
	kill_sigs := make(chan os.Signal, 1)
	signal.Notify(kill_sigs, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It works")
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bye Bye")
		bye_sigs <- true
	})

	server := &http.Server{Addr: "0.0.0.0:8081"}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	// Shutdown the server if: 1.Context ended 2.Receive signal from bye_sigs
	g.Go(func() error {
		select {
		case <-ctx.Done(): // Context ended by other co-routine

		case <-bye_sigs: // Receive signal from bye_sigs

		}
		return server.Shutdown(ctx)
	})

	g.Go(func() error {
		select {
		case <-ctx.Done(): // Context ended by other co-routine
			return nil
		case sig := <-kill_sigs: // OS kill signal
			fmt.Println("You killed me!")
			return errors.New(fmt.Sprint(sig))
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Program Exit")
}
