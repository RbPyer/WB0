package main

import (
	"log"
	"os"

	"github.com/RbPyer/WB0/internal/server"
	"github.com/RbPyer/WB0/internal/handler"
	"github.com/RbPyer/WB0/internal/repository"
	"github.com/RbPyer/WB0/internal/service"
	"github.com/RbPyer/WB0/internal/cache"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while loading env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("dbhost"),
		Port: viper.GetString("dbport"),
		Username: viper.GetString("dbuser"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("dbname"),
		SSLMode: viper.GetString("sslmode"),
	})

	if err != nil {
		log.Fatalf("error while initializing database: %s", err.Error())
	}
	cache := cache.NewCache()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(cache)

	srv := server.NewServer(viper.GetString("port"), handlers.InitRouting(), services, cache)
	srv.CacheLoad()
	if err := srv.Run(viper.GetString("sub")); err != nil {
		log.Fatalf("error while starting http-server: %s", err.Error())
	}
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}