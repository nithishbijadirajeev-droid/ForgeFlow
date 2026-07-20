package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/config"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/server"
)

func main() {

	config.Load()

	router := server.NewRouter()

	addr := fmt.Sprintf(
		"%s:%d",
		config.AppConfig.Server.Host,
		config.AppConfig.Server.Port,
	)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Println("ForgeFlow starting on", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctx)
}