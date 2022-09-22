package docker

import "github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"

func (docker *Docker) Sync(deployment *deployments.Deployment) error {
	docker.Delete(deployment.Name)

	for _, container := range deployment.Containers {
		err := docker.schedule(container, deployment.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (docker *Docker) Delete(deploymentName string) {
	containers, err := docker.listContainers(deploymentName)
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		docker.unschedule(docker.getContainerIdentifier(container, deploymentName))
	}
}

func (docker *Docker) List() []string {
	deployments := make([]string, 0)
	return deployments
}
