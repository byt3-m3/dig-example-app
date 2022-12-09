package handlers

import (
	"dig_practice/internal/models/cars"
	"dig_practice/internal/models/cars/repository"
	"fmt"
	"log"
	"net/http"
)

type CarHandler struct {
	Repo       repository.CarRepo
	RepoConfig *repository.CarRepoConfig
}

func (h *CarHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("invoke hit")
	fmt.Println("Repo config", h.RepoConfig)
	bmw := cars.Car{
		Make:  "bmw",
		Model: "5-series",
		Trim:  "m550i",
	}
	fmt.Println(h.Repo)

	if err := h.Repo.SaveCar(bmw); err != nil {
		log.Println("unable to save car")
	}
}

func NewCarHandler(cfg *repository.CarRepoConfig, carRepo repository.CarRepo) *CarHandler {
	return &CarHandler{Repo: carRepo, RepoConfig: cfg}
}
