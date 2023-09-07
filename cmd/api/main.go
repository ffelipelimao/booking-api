package main

import (
	"github.com/ffelipelimao/booking/internal/database"
	"github.com/ffelipelimao/booking/internal/repositories"
	create_place_usecase "github.com/ffelipelimao/booking/internal/usecases/create_place"

	webhandlers "github.com/ffelipelimao/booking/internal/web/handlers/create_place"
	webserver "github.com/ffelipelimao/booking/internal/web/server"
)

func main() {
	db := database.NewDatabase()
	defer database.Close(db)

	webserver := webserver.NewWebServer(":8080")

	placeRepository := repositories.NewPlaceRepository(db)
	createPlace := create_place_usecase.NewCreatePlaceUseCase(placeRepository)

	handlerCreatePlace := webhandlers.NewCreatePlaceHandler(createPlace)
	webserver.AddHandler("v1/places", handlerCreatePlace.Handle, "POST")

	webserver.Start()
}
