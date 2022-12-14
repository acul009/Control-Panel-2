package docker

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"

	"github.com/docker/docker/api/types"
	dockerContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	dockerErrors "github.com/docker/docker/errdefs"
	"github.com/docker/go-connections/nat"
)

func (docker *Docker) schedule(container *deployments.Container, deploymentName string, parameters []*deployments.Parameter) error {

	identifier := docker.getContainerIdentifier(container.Name, deploymentName)

	currentConfig, err := docker.cli.ContainerInspect(docker.ctx, identifier)
	if err != nil {
		fmt.Println(err)
	} else {
		currentHash := currentConfig.Config.Labels[CONTAINER_HASH]
		newHash := createContainerHash(container, parameters)
		if currentHash == newHash {
			//if hashes match, theres no need to recreate the container
			return nil
		}
	}

	err = docker.createContainer(container, true, deploymentName, parameters)

	switch err.(type) {

	case dockerErrors.ErrConflict:
		docker.deleteContainer(identifier)
		err = docker.createContainer(container, false, deploymentName, parameters)
	}

	if err != nil {
		return fmt.Errorf("error creating container %s: %w", container.Name, err)
	}

	err = docker.startContainer(identifier)

	if err != nil {
		return fmt.Errorf("error creating container %s: %w", container.Name, err)
	}

	return nil
}

func (docker *Docker) getContainerIdentifier(containerName string, deploymentName string) string {
	return docker.schedulerName + "-" + deploymentName + "-" + containerName
}

const MANGER_LABEL = LABEL_PREFIX + ".manager"
const DEPLOYMENT_LABEL = LABEL_PREFIX + ".deployment"
const CONTAINER_LABEL = LABEL_PREFIX + ".name"
const CONTAINER_HASH = LABEL_PREFIX + ".hash"

func (docker *Docker) createContainer(container *deployments.Container, pullImage bool, deploymentName string, parameters []*deployments.Parameter) error {
	var err error
	if pullImage {
		err = docker.pullImage(container.Image)
	}

	if err != nil {
		return fmt.Errorf("error requesting image: %w", err)
	}

	environment, err := docker.generateEnvParameterBindings(deploymentName, container.Parameters, parameters)

	if err != nil {
		return fmt.Errorf("error creating environment: %f", err)
	}

	bindList := docker.generateFileParameterBindings(deploymentName, container.Parameters, parameters)

	hash := createContainerHash(container, parameters)

	_, err = docker.cli.ContainerCreate(docker.ctx, &dockerContainer.Config{
		Image: container.Image,
		Labels: map[string]string{
			MANGER_LABEL:     docker.schedulerName,
			DEPLOYMENT_LABEL: deploymentName,
			CONTAINER_LABEL:  container.Name,
			CONTAINER_HASH:   hash,
		},
		Env: environment,
	},
		&dockerContainer.HostConfig{
			RestartPolicy: dockerContainer.RestartPolicy{
				Name: "always",
			},
			Tmpfs: map[string]string{
				"/tmp": "rw",
			},
			PortBindings: docker.createPortMap(container),
			//ReadonlyRootfs: true,
			Binds: bindList,
		},
		nil, nil,
		docker.getContainerIdentifier(container.Name, deploymentName),
	)

	return err
}

func createContainerHash(container *deployments.Container, parameters []*deployments.Parameter) string {
	sb := strings.Builder{}

	paramValues := make(map[string]string)

	for _, param := range parameters {
		paramValues[param.Name] = param.Value
	}

	for _, param := range container.Parameters {
		sb.WriteString(paramValues[*param.Name])
		sb.WriteString("\n")
	}

	config, _ := json.Marshal(container)

	sb.Write(config)
	hasher := crypto.MD5.New()
	hasher.Write([]byte(sb.String()))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (docker *Docker) createPortMap(container *deployments.Container) nat.PortMap {
	var ports nat.PortMap = make(nat.PortMap, len(container.Ports))
	for _, portMapping := range container.Ports {
		ports[nat.Port(fmt.Sprintf("%v/%s", portMapping.Container, portMapping.Protocol))] = []nat.PortBinding{{
			HostPort: fmt.Sprintf("%v", portMapping.Host),
		}}
	}
	return ports
}

func (docker *Docker) pullImage(image string) error {
	reader, err := docker.cli.ImagePull(docker.ctx, image, types.ImagePullOptions{})
	defer reader.Close()
	var readerError error = nil
	var buf []byte = make([]byte, 1024)
	for ; readerError != io.EOF; _, readerError = reader.Read(buf) {
	}

	if err != nil {
		return fmt.Errorf("error pulling image: %w", err)
	}
	return nil
}

func (docker *Docker) deleteContainer(identifier string) error {
	err := docker.cli.ContainerRemove(docker.ctx, identifier, types.ContainerRemoveOptions{
		Force: true,
	})

	if err != nil {
		return fmt.Errorf("error deleting container: %w", err)
	}

	return nil
}

func (docker *Docker) startContainer(identifier string) error {
	err := docker.cli.ContainerStart(docker.ctx, identifier, types.ContainerStartOptions{})

	if err != nil {
		return fmt.Errorf("error starting container: %w", err)
	}

	return nil
}

func (docker *Docker) unschedule(identifier string) error {
	err := docker.deleteContainer(identifier)
	if err != nil {
		return fmt.Errorf("error unscheduling deployment: %w", err)
	}
	return nil
}

func (docker *Docker) listContainers(deployment string) ([]string, error) {
	filters := filters.NewArgs(
		filters.KeyValuePair{
			Key:   "label",
			Value: MANGER_LABEL + "=" + docker.schedulerName,
		},
		filters.KeyValuePair{
			Key:   "label",
			Value: DEPLOYMENT_LABEL + "=" + deployment,
		},
	)
	return docker.listContainersForFilter(filters)
}

func (docker *Docker) listContainersForFilter(filterList filters.Args) ([]string, error) {
	var dockerContainers []types.Container
	var err error

	dockerContainers, err = docker.cli.ContainerList(docker.ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterList,
	})

	if err != nil {
		return make([]string, 0), fmt.Errorf("error loading Deployments: %w", err)
	}

	var containers []string = make([]string, 0, len(dockerContainers))

	for _, container := range dockerContainers {

		if container.Labels[MANGER_LABEL] != docker.schedulerName {
			continue
		}

		containers = append(containers, container.Labels[CONTAINER_LABEL])
	}
	return containers, nil
}
