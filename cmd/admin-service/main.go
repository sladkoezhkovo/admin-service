package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	api "github.com/sladkoezhkovo/admin-service/api/admin"
	"github.com/sladkoezhkovo/admin-service/internal/config"
	"github.com/sladkoezhkovo/admin-service/internal/handler"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	cityrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/city-repository"
	ctrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/ct-repository"
	districtrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/district-repository"
	packagingrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/packaging-repository"
	ptrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/property-repository"
	unitrepository "github.com/sladkoezhkovo/admin-service/internal/repository/pg/unit-repository"
	cityservice "github.com/sladkoezhkovo/admin-service/internal/service/city-service"
	ctservice "github.com/sladkoezhkovo/admin-service/internal/service/ct-service"
	districtservice "github.com/sladkoezhkovo/admin-service/internal/service/district-service"
	packagingservice "github.com/sladkoezhkovo/admin-service/internal/service/packaging-service"
	ptservice "github.com/sladkoezhkovo/admin-service/internal/service/pt-service"
	unitservice "github.com/sladkoezhkovo/admin-service/internal/service/unit-service"
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

	fmt.Println(cfg)

	fmt.Printf("db init")
	db, err := pg.Setup(&cfg.Pg)
	if err != nil {
		panic(fmt.Sprintf("setup database: %s", err.Error()))
	}

	cityRepository := cityrepository.New(db)
	unitRepository := unitrepository.New(db)
	districtRepository := districtrepository.New(db)
	packagingRepository := packagingrepository.New(db)
	ptRepository := ptrepository.New(db)
	ctRepository := ctrepository.New(db)

	cityService := cityservice.New(cityRepository)
	unitService := unitservice.New(unitRepository)
	districtService := districtservice.New(districtRepository)
	packagingService := packagingservice.New(packagingRepository)
	ptService := ptservice.New(ptRepository)
	ctService := ctservice.New(ctRepository)

	server := grpc.NewServer()
	adminServiceHandler := handler.New(
		cityService,
		districtService,
		packagingService,
		unitService,
		ptService,
		ctService,
	)

	api.RegisterAdminServiceServer(server, adminServiceHandler)

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
