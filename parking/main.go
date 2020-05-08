package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"parking/src"
	"syscall"
	"time"
)

func main() {
	//addr := os.Getenv("LISTEN_ADDR")
	httpServer := &http.Server{
		Addr:    ":" + "10002",
		Handler: src.Mux(),
	}
	go func() {
		kill := make(chan os.Signal, 1)
		signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
		<-kill
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		httpServer.Shutdown(ctx)
		cancel()
	}()
	fmt.Println("zzzz")
	//fmt.Println("start", zap.String("listen", addr))
	err := httpServer.ListenAndServe()
	fmt.Println("finish", zap.String("message", err.Error()))
}
