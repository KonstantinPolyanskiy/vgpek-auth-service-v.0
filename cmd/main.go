package main

import (
	"awesomeProject5"
	handlers2 "awesomeProject5/pkg/handlers"
	"awesomeProject5/pkg/repository"
	"awesomeProject5/pkg/service"
	_ "database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config - %v", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка загрузки .env - %v", err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("test-db.host"),
		Port:     viper.GetString("test-db.port"),
		Username: viper.GetString("test-db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("test-db.dbname"),
		SSLMode:  viper.GetString("test-db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка подключения к бд - %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers2.NewHandler(services)

	srv := new(awesomeProject5.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error - %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
