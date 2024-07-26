package main

import (
	"learn_rabbitmq/model/appModel"
	"learn_rabbitmq/router"
)

func main() {
	appModel.GenerateAppModel()
	router.SetupRouter()
}
