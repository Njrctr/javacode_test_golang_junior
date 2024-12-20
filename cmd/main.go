package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	handler "github.com/Njrctr/javacode_test_golang_junior/pkg/handlers"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	pg_rep "github.com/Njrctr/javacode_test_golang_junior/pkg/repository/postgres"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Wallet API
// @version 1.0
// @description API Server for Wallet

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка получения переменных окружения: %s", err.Error())
	}
	db, err := pg_rep.NewDB(pg_rep.Config{
		// Host: viper.GetString("db.host"),
		Host: "localhost",
		// Port: viper.GetString("db.port"),
		Port:     "5436",
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка инициализации Базы данных: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(models.Server)
	logrus.Printf("Попытка запуска сервера на порту %s", viper.GetString("port"))

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Walletter APP Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Walletter App Stoped")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
