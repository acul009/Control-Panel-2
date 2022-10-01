package database

import (
	"fmt"

	"github.com/acul009/control-panel-2/src/api/deployments/database/filesystemdb"
	"github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"
	"github.com/acul009/control-panel-2/src/api/deployments/scheduler"
	"github.com/spf13/viper"
)

type Database interface {
	Save(deployment *deployments.Deployment) error
	List() []string
	Load(deploymentName string) (*deployments.Deployment, error)
	Delete(deploymentName string) error
}

var db Database = nil

func GetDatabase() Database {
	fmt.Println("requesting Database")
	if db == nil {
		var provider Database
		viper.AutomaticEnv()
		viper.SetDefault("database.provider", "fs")

		switch viper.GetString("database.provider") {
		case "fs":
			provider = filesystemdb.CreateFilesystemDB()
			break
		default:
			panic("unknown database provider")
		}
		db = databaseHook{
			provider,
		}
	}
	return db
}

type databaseHook struct {
	Database
}

func (hook databaseHook) Save(deployment *deployments.Deployment) error {
	sched := scheduler.GetScheduler()
	err := sched.Sync(deployment)
	if err != nil {
		return err
	}
	err = hook.Database.Save(deployment)
	return err
}

func (hook databaseHook) List() []string {
	return hook.Database.List()
}

func (hook databaseHook) Load(deploymentName string) (*deployments.Deployment, error) {
	return hook.Database.Load(deploymentName)
}

func (hook databaseHook) Delete(deploymentName string) error {
	sched := scheduler.GetScheduler()
	sched.Delete(deploymentName)
	return hook.Database.Delete(deploymentName)
}
