package webserver

import (
	"context"
	"dig_practice/internal/application/car_manager/webserver/handlers"
	"fmt"
	"log"
	"net/http"
)

type Webserver struct {
	carHandler *handlers.CarHandler
}

type CarGetter interface {
	GetCars() error
}

func New(carHandler *handlers.CarHandler, carGetter CarGetter) *Webserver {
	return &Webserver{carHandler: carHandler}
}

func (w *Webserver) Run(ctx context.Context) error {
	fmt.Println("starting Webserver")

	router := http.NewServeMux()

	router.HandleFunc("/api/v1/cars", w.carHandler.HandleGet)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal("unable to  run Webserver")
	}

	return nil
}
