package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/andreipimenov/carservice/driver"
	"github.com/andreipimenov/carservice/interactor"
	"github.com/andreipimenov/carservice/storage"
)

func main() {

	cfgPtr := flag.String("config", "", "path to the configuration file")
	flag.Parse()

	cfgFile := "config.json"
	if *cfgPtr != "" {
		cfgFile = *cfgPtr
	}

	cfg, err := NewConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	mongoDriver, err := driver.NewMongo(cfg.MongoURL, cfg.MongoDB)
	if err != nil {
		log.Fatal(err)
	}

	carStorage := storage.NewCar(mongoDriver, "car")
	carInteractor := interactor.NewCar(carStorage)

	http.Handle("/", NotFoundHandler())
	http.Handle("/api/ping", PingHandler())
	http.Handle("/api/cars", CarsHandler(carInteractor))

	log.Printf("Server is listening on port %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
