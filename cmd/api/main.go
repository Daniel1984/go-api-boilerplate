package main

import (
	"log"

	"github.com/boilerplate/cmd/api/router"
	"github.com/boilerplate/pkg/application"
	"github.com/boilerplate/pkg/exithandler"
	"github.com/boilerplate/pkg/logger"
	"github.com/boilerplate/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			log.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
