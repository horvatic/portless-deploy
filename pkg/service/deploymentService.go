package service

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/horvatic/portless-deploy/pkg/store"
)

type DeploymentService interface {
	StartDeployment()
}

type deploymentService struct {
	store store.Store
}

func BuildDeploymentService(store store.Store) DeploymentService {
	return &deploymentService{
		store: store,
	}
}

func (d *deploymentService) StartDeployment() {
	deployments := d.store.GetAllDeployments()
	for _, deployment := range deployments {
		d.store.DeleteDeployment(deployment.DeploymentId)
	}
	for _, deployment := range deployments {
		fmt.Println(deployment.RepoUri)
		fmt.Println("/deployments/work/"+deployment.DeploymentId)
		branch := strings.Replace(deployment.BranchName, "refs/heads/", "", -1)
		cloneCmd := exec.Command("git", "clone", "--branch", branch, deployment.RepoUri, "/deployments/work/"+deployment.DeploymentId)
		cloneErr := cloneCmd.Run()
		if cloneErr != nil {
			fmt.Println(cloneErr)
			continue
		}


		permissionCmd := exec.Command("chmod", "+x", "/deployments/work/"+deployment.DeploymentId+"/deploy/scripts/"+deployment.TargetScript)
                permissionErr := permissionCmd.Run()
                if permissionErr != nil {
                        fmt.Println(permissionErr)
                        continue
                }

		deployCmd := exec.Command("bash", "-c", "/deployments/work/"+deployment.DeploymentId+"/deploy/scripts/"+deployment.TargetScript + " " + deployment.GitShortSha + " " + deployment.DeploymentId + " " + deployment.Env)
                deployErr := deployCmd.Run()
		if deployErr != nil {
                        fmt.Println(deployErr)
                        continue
                }

		removeCmd := exec.Command("rm", "-rf", "/deployments/work/"+deployment.DeploymentId)
                removeErr := removeCmd.Run()
                if removeErr != nil {
                        fmt.Println(removeErr)
                        continue
                }
	}
}

