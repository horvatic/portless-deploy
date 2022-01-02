package main

import (
	"fmt"
	"os"
	"time"

	"github.com/horvatic/portless-deploy/pkg/service"
	"github.com/horvatic/portless-deploy/pkg/store"
)

func main() {

	connectionString := os.Args[1]
	database := os.Args[2]
	collection := os.Args[3]

	for {
		store, dbClient, dbContext, err := store.BuildMongoDeploymentStore(connectionString, database, collection)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		s := service.BuildDeploymentService(store)

		s.StartDeployment()

		dbClient.Disconnect(dbContext)
		time.Sleep(5 * time.Minute)
	}
}
