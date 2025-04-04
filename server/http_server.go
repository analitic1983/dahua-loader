package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"koshmin/dahua-loader/config"
	"koshmin/dahua-loader/server/pages"
	"log"
	"net/http"
	"time"
)

func HTTPServerStart(ctx context.Context) {

	const CLOSE_TIMEOUT = 15 * time.Second

	// Wait for finish signal
	r := gin.Default()

	// Routes
	r.GET("/", pages.Index)
	r.GET("/ping", pages.Ping)

	// Start on port 8080
	listenOn := config.AppConfig.HttpServer.Host + ":" + config.AppConfig.HttpServer.Port
	log.Println("Http server listen on ", listenOn)

	srv := &http.Server{
		Addr:    listenOn,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Http server failed: ", err)
			panic(err.Error())
		}
	}()

	// Wait for close signal
	select {
	case <-ctx.Done():
	}

	// Create context for graceful finish
	ctx, cancel := context.WithTimeout(context.Background(), CLOSE_TIMEOUT)
	defer cancel()

	// Shutdown server with timeout
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Http server forced to shutdown: %s\n", err)
	}
	log.Println("Http server closed")
}
