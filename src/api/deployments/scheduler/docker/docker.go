package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/spf13/viper"
)

type Docker struct {
	cli           *client.Client
	ctx           context.Context
	schedulerName string
	bindLocation  string
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

	viper.SetDefault("scheduler.docker.bindlocation", "/var/control-panel-2/bind")

	var scheduler *Docker = &Docker{
		cli:           cli,
		schedulerName: schedulerName,
		ctx:           context.Background(),
		bindLocation:  viper.GetString("scheduler.docker.bindlocation") + "/" + schedulerName,
	}

	viper.Debug()

	fmt.Println("binding to:" + scheduler.bindLocation)

	return scheduler, nil
}
