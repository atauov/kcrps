package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/atauov/kcrps"
	"github.com/atauov/kcrps/pkg/handler"
	"github.com/atauov/kcrps/pkg/repository"
	"github.com/atauov/kcrps/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title kaspi custom API service
// @version 0.1
// @description API server for custom remote payment service in kaspi mobile POS

// @host 145.249.246.27
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Cant load env variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(kcrps.Server)

	go func() {
		if err = srv.Run(viper.GetString("port"), handlers.InitRoutes(), "key", "key2"); err != nil {
			logrus.Fatalf("Can not to run server, %s", err.Error())
		}
	}()

	logrus.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server shutting down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
