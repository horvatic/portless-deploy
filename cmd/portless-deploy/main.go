package main

import (
	"os"
	"strconv"
	"time"

	"github.com/horvatic/portless-deploy/pkg/service"
	"github.com/horvatic/portless-deploy/pkg/store"
)

func main() {

	connectionString := os.Args[1]
	database := os.Args[2]
	collection := os.Args[3]
	sleepInSeconds, convertErr := strconv.Atoi(os.Args[4])
	if convertErr != nil {
		return
	}

	for {
		store, dbClient, dbContext, err := store.BuildMongoDeploymentStore(connectionString, database, collection)
		if err != nil {
			return
		}

		s := service.BuildDeploymentService(store)

		s.StartDeployment()

		dbClient.Disconnect(dbContext)
		time.Sleep(time.Duration(sleepInSeconds) * time.Second)
	}
}
