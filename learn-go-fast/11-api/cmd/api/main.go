package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/cankurttekin/gotutorial/11-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("starting awesome go server")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

}
	

