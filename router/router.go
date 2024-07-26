package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"learn_rabbitmq/endpoint"
	"learn_rabbitmq/model/errorModel"
	"log"
	"net/http"
	"os"
)

func SetupRouter() {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf(`/publish`), endpoint.RabbitPublishEndpoint).Methods("POST")
	r.HandleFunc(fmt.Sprintf(`/consume`), endpoint.RabbitConsumeEndpoint).Methods("GET")
	log.Printf(`server running on port %s ...`, os.Getenv("APP_PORT"))
	err := http.ListenAndServe(fmt.Sprintf(`:%s`, os.Getenv("APP_PORT")), r)
	if err != nil {
		errorModel.FailOnErr(err, "Failed create http server...")
	}
}
