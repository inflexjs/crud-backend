package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/inflexjs/crud-backend"
	"github.com/inflexjs/crud-backend/internal/handler"
	"github.com/inflexjs/crud-backend/internal/service"
	"github.com/inflexjs/crud-backend/internal/storage"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация .yml конфига
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing .yml config: %s", err.Error())
	}

	// Инициализация .env конфига
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing .env config: %s", err.Error())
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Инициализация зависимостей
	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	// Инициализация net/http сервера
	srv := new(crud.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
