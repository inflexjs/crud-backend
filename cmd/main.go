package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/inflexjs/crud-backend"
	"github.com/inflexjs/crud-backend/internal/handler"
	"github.com/inflexjs/crud-backend/internal/service"
	"github.com/inflexjs/crud-backend/internal/storage"
	"github.com/spf13/viper"
)

func main() {
	// Определение формата JSON для logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация .yml конфига
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing .yml config: %s", err.Error())
	}

	// Инициализация .env конфига
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing .env config: %s", err.Error())
	}

	// Инициализация database
	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Инициализация зависимостей
	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	// Gin mode - release or debug
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Инициализация net/http сервера
	srv := new(crud.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
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
