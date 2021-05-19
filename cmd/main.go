package main

import (
	"log"
	"os"

	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/DanYesmagulov/go-phone-book/pkg/handler"
	"github.com/DanYesmagulov/go-phone-book/pkg/repository"
	"github.com/DanYesmagulov/go-phone-book/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка во время инициализации переменных окружения: %s", err.Error())
	}
	db, err := repository.NewPgDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("Ошибка во время инициализации бд: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(phonebook.Server)
	if err := server.Run("9000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка во время запуска сервера: %s", err.Error())
	}
}
