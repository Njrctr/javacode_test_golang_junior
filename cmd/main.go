package main

import (
	"context"
	"fmt"
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

	config, err := initConfig()
	if err != nil {
		logrus.Fatalf("Ошибка инициализации Конфига: %s", err.Error())
	}

	fmt.Println(config)

	db, err := pg_rep.NewDB(pg_rep.Config{
		Host:     config["DB_HOST"],
		Port:     config["DB_PORT"],
		Username: config["DB_USERNAME"],
		DBName:   config["DB_DBNAME"],
		SSLMode:  config["DB_SSLMODE"],
		Password: config["DB_PASSWORD"],
	})
	if err != nil {
		logrus.Fatalf("Ошибка инициализации Базы данных: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(models.Server)
	logrus.Printf("Попытка запуска сервера на порту %s", config["APP_PORT"])

	go func() {
		if err := server.Run(config["APP_PORT"], handlers.InitRouters()); err != nil && err != http.ErrServerClosed {
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

func initConfig() (map[string]string, error) {
	if err := godotenv.Load("config.env"); err != nil {
		logrus.Fatalf("Ошибка получения переменных окружения: %s", err.Error())
	}

	config := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_DBNAME":   os.Getenv("DB_DBNAME"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"APP_PORT":    os.Getenv("APP_PORT"),
	}
	return config, nil
}
