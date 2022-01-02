package main

import (
	"fmt"
	"os"

	"github.com/horvatic/portless-deploy/pkg/service"
	"github.com/horvatic/portless-deploy/pkg/store"
)

func main() {
	store, dbClient, dbContext, err := store.BuildMongoDeploymentStore(os.Getenv("MONGO_CONNECTION_STRING"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_COLLECTION"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := service.BuildDeploymentService(store)

	s.StartDeployment()

	fmt.Println("Server Stopped")
	dbClient.Disconnect(dbContext)
	fmt.Println("Db Disconnected")
}
