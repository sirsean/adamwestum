package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirsean/adamwestum/api"
	"github.com/sirsean/adamwestum/config"
	"github.com/sirsean/adamwestum/controller"
	"github.com/sirsean/adamwestum/service"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting Up")

	adamWestFile, err := os.Open(fmt.Sprintf("%s/quotes/adamwest", config.Get().Host.Path))
	if err != nil {
		log.Fatal("Could not read adamwest quotes")
	}
	defer adamWestFile.Close()
	service.Load(adamWestFile)

	router := mux.NewRouter()
	router.HandleFunc("/", controller.Index).Methods("GET")

	router.HandleFunc("/api/lines/{number}", api.Lines).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(fmt.Sprintf("%s/static/", config.Get().Host.Path))))
	http.Handle("/", router)

	port := config.Get().Host.Port
	log.Fatal(http.ListenAndServe(port, nil))
}
