package main

import (
	"context"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"koshmin/dahua-loader/config"
	"koshmin/dahua-loader/database"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	config.Init()

	// InitSQL database
	if err := database.InitSQL(); err != nil {
		log.Fatal("Failed db connect: ", err)
		panic(err.Error())
	}
	if err := database.InitRedis(); err != nil {
		log.Fatal("Failed redis connect: ", err)
		panic(err.Error())
	}

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Chanel for os signals
	signalChain := make(chan os.Signal, 1)
	signal.Notify(signalChain, syscall.SIGINT, syscall.SIGTERM) // Subscribe on SIGINT (Ctrl+C) и SIGTERM (terminate process)

	// Start http server
	wg.Add(1)
	go func() {
		defer wg.Done()
		HTTPServerStart(ctx)
	}()

	signalInfo := <-signalChain
	log.Println("Terminate signal received: " + signalInfo.String())
	cancel() // Send notifications to close
	wg.Wait()
}
