# week 03 作业

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

```go
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
			case syscall.SIGTERM:
			case syscall.SIGINT:
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
```

## comment

基本正确，信号量的处理需要针对：syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT。
