package main

import (
	"flag"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

var configPath string
var dotenv bool

func init() {
	flag.StringVar(&configPath, "config", "configs/.yml", "config path")
	flag.BoolVar(&dotenv, "dotenv", false, "turn on dotenv")
}

func main() {
	flag.Parse()

	server := grpc.NewServer()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	<-stopChan

	server.GracefulStop()
}
