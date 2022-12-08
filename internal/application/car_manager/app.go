package car_manager

import (
	"context"
	"dig_practice/internal/application/car_manager/webserver"
	"dig_practice/internal/application/car_manager/worker"
	"dig_practice/internal/models/cars/repository"
	"go.uber.org/dig"
	"log"
	"os"
)

type icomponent interface {
	Run(ctx context.Context) error
}

type app struct {
}

func New() *app {
	return &app{}
}

func NewRepoConfig() *repository.CarRepoConfig {
	cfg := &repository.CarRepoConfig{}
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBTable = os.Getenv("DB_TABLE")
	return cfg
}

func (a app) Run(ctx context.Context, component string) error {

	c := dig.New()

	if err := c.Provide(NewRepoConfig); err != nil {
		log.Println("unable to add config dep")
	}
	if err := c.Provide(repository.NewCarRepo, dig.As(new(repository.CarRepo)), dig.As(new(webserver.CarGetter))); err != nil {
		log.Println("unable to add repo dep")
	}
	if err := c.Provide(webserver.New); err != nil {
		log.Println("could not provide webserver")
	}

	//if err := c.Provide(handlers.NewCarHandler); err != nil {
	//	log.Println("could not provide handler")
	//}

	switch component {
	case "worker":
		if err := worker.New().Run(ctx); err != nil {
			log.Println("unable to run worker")
			return err
		}

	case "webserver":

		if err := c.Invoke(func(ws *webserver.Webserver) {

			if err := ws.Run(ctx); err != nil {
				log.Println("unable to run webserver")

			}

		}); err != nil {
			log.Println(err)
		}

	}

	return nil

}
