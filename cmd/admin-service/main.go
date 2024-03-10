package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sladkoezhkovo/admin-service/api"
	"github.com/sladkoezhkovo/admin-service/internal/config"
	handler "github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	cityrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/city-repository"
	cityservice "github.com/sladkoezhkovo/admin-service/internal/service/city-service"
	"github.com/sladkoezhkovo/lib"
	"google.golang.org/grpc"
	"net"
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

	if dotenv {
		fmt.Printf("loading dotenv...\n")
		if err := godotenv.Load(".env.pg"); err != nil {
			panic(fmt.Errorf("godotenv.Load: %s", err))
		}
	}

	var cfg config.Config

	if err := lib.SetupConfig(configPath, &cfg); err != nil {
		panic(fmt.Errorf("cannot read config: %s", err))
	}

	fmt.Printf("db init")
	db, err := pg.Setup(&cfg.Pg)
	if err != nil {
		panic(fmt.Sprintf("setup database: %s", err.Error()))
	}

	fmt.Printf("\t+\ncity repository init")
	cityRepository := cityrepository.New(db)

	fmt.Printf("\t+\ncity service init")
	cityService := cityservice.New(cityRepository)
	fmt.Printf("\t+\n")

	server := grpc.NewServer()
	adapter := handler.New(
		cityService,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	api.RegisterAdminServiceServer(server, adapter)

	go func(s *grpc.Server, cfg *config.AppConfig) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
		if err != nil {
			panic(fmt.Errorf("cannot bind port %d", cfg.Port))
		}
		fmt.Printf("\nServer started on %d port\n\n", cfg.Port)
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}(server, &cfg.App)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	<-stopChan

	server.GracefulStop()
}
