package docker

import "github.com/acul009/control-mono/api/deployments/gen/deployments"

func (docker *Docker) Sync(deployment *deployments.Deployment) error {
	containers, err := docker.listContainers(deployment.Name)
	if err != nil {
		return err
	}

	activeContainers := make(map[string]struct{}, len(containers))

	for _, container := range containers {
		activeContainers[container] = struct{}{}
	}

	for _, container := range deployment.Containers {
		delete(activeContainers, container.Name)
		err = docker.schedule(container, deployment.Name)
		if err != nil {
			return err
		}
	}

	for orphanContainer := range activeContainers {
		err = docker.unschedule(docker.getContainerIdentifier(orphanContainer, deployment.Name))
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
