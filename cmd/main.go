package main

import (
	"context"
	"dig_practice/internal/application/car_manager"
	"log"
)

type App interface {
	Run(ctx context.Context, component string) error
}

func main() {
	ctx := context.Background()

	if err := car_manager.New().Run(ctx, "webserver"); err != nil {
		log.Fatal("unable to run app")
	}
}
