package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

type Docker struct {
	cli           *client.Client
	ctx           context.Context
	schedulerName string
}

const LABEL_PREFIX = "de.it-woelfchen.controlPanel"

func CreateDockerScheduler(schedulerName string) (*Docker, error) {
	var cli *client.Client
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return nil, fmt.Errorf("error creating docker scheduler: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("error creating docker scheduler: %w", err)
	}

	var scheduler *Docker = &Docker{
		cli:           cli,
		schedulerName: schedulerName,
		ctx:           context.Background(),
	}

	return scheduler, nil
}
