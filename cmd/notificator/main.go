package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BarTar213/notificator/api"
	"github.com/BarTar213/notificator/config"
	"github.com/BarTar213/notificator/email"
	"github.com/BarTar213/notificator/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig("notificator.yml")
	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Printf("%+v\n", conf)

	if conf.Api.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	postgres, err := storage.NewPostgres(&conf.Postgres)
	if err != nil {
		logger.Fatalf("new postgres: %s", err)
	}

	emailCli, err := email.New(&conf.Mail)
	if err != nil {
		logger.Fatalf("email client: %s", err)
	}

	a := api.NewApi(
		api.WithConfig(conf),
		api.WithLogger(logger),
		api.WithStorage(postgres),
		api.WithEmailClient(emailCli),
	)

	go a.Run()
	logger.Print("started app")

	shutDownSignal := make(chan os.Signal)
	signal.Notify(shutDownSignal, syscall.SIGINT, syscall.SIGTERM)

	<-shutDownSignal
	logger.Print("exited from app")
}
