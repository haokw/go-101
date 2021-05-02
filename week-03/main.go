package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	serveRun := func(ctx context.Context) error {
		g, ctx := errgroup.WithContext(ctx)
		notify := []os.Signal{syscall.SIGQUIT}
		sgc := make(chan os.Signal)
		stop := make(chan struct{})
		g.Go(func() error {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
				fmt.Fprintln(resp, "Hello, Go")
			})
			return serve("0.0.0.0:8080", mux, notify, sgc, stop)
		})
		g.Go(func() error {
			return serve("127.0.0.1:8001", http.DefaultServeMux, notify, sgc, stop)
		})
		if err := g.Wait(); err != nil {
			return err
		}
		return nil
	}

	err := serveRun(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func main02() {
// 	done := make(chan error, 2)
// 	stop := make(chan struct{})
// 	go func() {
// 		done <- serveDebug(stop)
// 	}()
// 	go func() {
// 		done <- serveApp(stop)
// 	}()

// 	var stopped bool
// 	for i := 0; i < cap(done); i++ {
// 		if err := <-done; err != nil {
// 			fmt.Printf("error: %v", err)
// 		}
// 		if !stopped {
// 			stopped = true
// 			close(stop)
// 		}
// 	}
// }

func serve(addr string, handler http.Handler, notify []os.Signal, sgc chan os.Signal, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	signal.Notify(sgc, notify...)

	go func() {
		for sg := range sgc {
			switch sg {
			case syscall.SIGQUIT:
				s.Shutdown(context.Background())
			default:
			}
		}
	}()

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(notify []os.Signal, sgc chan os.Signal, stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, Go")
	})
	return serve("0.0.0.0:8080", mux, notify, sgc, stop)
}

func serveDebug(notify []os.Signal, sgc chan os.Signal, stop <-chan struct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, notify, sgc, stop)
}
