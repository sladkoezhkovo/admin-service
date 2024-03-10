package main

import (
	"flag"
	"fmt"
	"github.com/sladkoezhkovo/admin-service/internal/config"
	"github.com/sladkoezhkovo/lib"
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

	var cfg config.Config

	if err := lib.SetupConfig(configPath, &cfg); err != nil {
		panic(fmt.Errorf("cannot read config: %s", err))
	}

	server := grpc.NewServer()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	<-stopChan

	server.GracefulStop()
}
