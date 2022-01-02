package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deployment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	DeploymentId   string             `bson:"deploymentId"`
	RepoUri        string             `bson:"repoUri"`
	DeploymentName string             `bson:"deploymentName"`
	Env            string             `bson:"env"`
	TargetScript   string             `bson:"targetScript"`
	GitSha         string             `bson:"gitSha"`
	BranchName     string             `bson:"branchName"`
        GitShortSha    string             `bson:"gitShortSha"`
}
