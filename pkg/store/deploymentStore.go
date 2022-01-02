package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/horvatic/portless-deploy/pkg/model"
)

type Store interface {
	GetAllDeployments() []model.Deployment
	DeleteDeployment(projectId string) error
}

type mongoDeploymentStore struct {
	client     *mongo.Client
	context    context.Context
	database   string
	collection string
}

func BuildMongoDeploymentStore(connectionString string, database string, collection string) (Store, *mongo.Client, context.Context, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return &mongoDeploymentStore{
		client:     client,
		database:   database,
		collection: collection,
	}, client, ctx, nil
}

func (m *mongoDeploymentStore) getDeploymentCollection() *mongo.Collection {
	db := m.client.Database(m.database)
	projects := db.Collection(m.collection)
	return projects
}

func (m *mongoDeploymentStore) GetAllDeployments() []model.Deployment {
	deploymentsCollection := m.getDeploymentCollection()
	cursor, err := deploymentsCollection.Find(m.context, bson.M{})
	if err != nil {
		return nil
	}
	var deployments []model.Deployment
	defer cursor.Close(m.context)
	for cursor.Next(m.context) {
		var deployment model.Deployment
		if err = cursor.Decode(&deployment); err != nil {
			return nil
		}
		deployments = append(deployments, deployment)
	}
	return deployments
}

func (m *mongoDeploymentStore) DeleteDeployment(deploymentId string) error {
	deploymentsCollection := m.getDeploymentCollection()
	_, err := deploymentsCollection.DeleteOne(m.context, bson.M{"deploymentId": deploymentId})
	if err != nil {
		return err
	}
	return nil
}
