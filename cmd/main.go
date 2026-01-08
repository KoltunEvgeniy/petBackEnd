package main

import (
	"meawby/internal/handler"
	"meawby/internal/repository"
	"meawby/internal/server"
	"meawby/internal/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	db, err := repository.NewConnectPostgres(repository.Config{
		Host:    "localhost",
		Port:    "5432",
		Dbname:  "meawby",
		Sslmode: "disable",
	})
	if err != nil {
		logrus.Fatal(err)
		return
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	srv := new(server.Server)
	if err := srv.Run("8000", handler.InitRouts()); err != nil {
		logrus.Fatal(err)
		return
	}

}
