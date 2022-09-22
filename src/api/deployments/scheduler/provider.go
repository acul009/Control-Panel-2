package scheduler

import (
	"github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"
	"github.com/acul009/control-panel-2/src/api/deployments/scheduler/docker"
	"github.com/spf13/viper"
)

type Scheduler interface {
	Sync(*deployments.Deployment) error
	Delete(deploymentName string)
	List() []string
}

var scheduler Scheduler = nil

func GetScheduler() Scheduler {
	if scheduler == nil {
		var err error
		viper.AutomaticEnv()
		viper.SetDefault("scheduler.provider", "docker")

		switch viper.GetString("scheduler.provider") {
		case "docker":
			viper.SetDefault("scheduler.docker.name", "controlpanel")
			scheduler, err = docker.CreateDockerScheduler(viper.GetString("scheduler.docker.name"))
			break
		default:
			panic("unknown database provider")
		}
		if err != nil {
			panic(err)
		}
	}
	return scheduler
}

//Need to put this elsewhere -> import cycle

// var db = database.GetDatabase()

// func FullSync() {
// 	deployments := db.List()
// 	scheduler := GetScheduler()
// 	active := scheduler.List()

// 	var activeSet map[string]struct{} = make(map[string]struct{}, len(active))
// 	for _, deploymentName := range active {
// 		activeSet[deploymentName] = struct{}{}
// 	}

// 	for _, deployment := range deployments {
// 		delete(activeSet, deployment)
// 		manifest, _ := db.Load(deployment)
// 		scheduler.Sync(manifest)
// 	}

// 	for deployment := range activeSet {
// 		scheduler.Delete(deployment)
// 	}
// }
