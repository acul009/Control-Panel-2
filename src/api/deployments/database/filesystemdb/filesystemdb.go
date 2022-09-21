package filesystemdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"os"

	dbErrors "github.com/acul009/control-mono/api/deployments/database/errors"
	"github.com/acul009/control-mono/api/deployments/gen/deployments"
	"github.com/spf13/viper"
)

type FilesystemDB struct {
	path string
}

func CreateFilesystemDB() *FilesystemDB {
	viper.SetDefault("database.fs.location", "/home/acul/control-mono/db/")
	db := &FilesystemDB{
		path: viper.GetString("database.fs.location"),
	}
	fmt.Println(viper.GetString("database.fs.location"))
	db.createLocation()
	return db
}

func (db FilesystemDB) createLocation() {
	err := os.MkdirAll(db.path, 0700)
	if err != nil {
		panic(err)
	}
}

func (db FilesystemDB) getFilePath(name string) string {
	return db.path + name + ".json"
}

func (db FilesystemDB) Save(deployment *deployments.Deployment) error {
	filename := db.getFilePath(deployment.Name)
	bytes, err := json.Marshal(deployment)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filename, bytes, 0700)
	return nil
}

func (db FilesystemDB) List() []string {
	files, err := ioutil.ReadDir(db.path)
	if err != nil {
		panic(err)
	}
	var list []string = make([]string, 0, len(files))

	for _, file := range files {
		name := file.Name()
		list = append(list, name[0:len(name)-5])
	}
	return list
}

func (db FilesystemDB) Load(name string) (*deployments.Deployment, error) {
	filename := db.getFilePath(name)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, dbErrors.MakeNotFoundError(name)
		}
		panic(err)
	}

	var bytes []byte
	bytes, err = io.ReadAll(file)

	deployment := &deployments.Deployment{}
	json.Unmarshal(bytes, deployment)
	return deployment, nil
}

func (db FilesystemDB) Delete(name string) error {
	err := os.Remove(db.getFilePath(name))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return dbErrors.MakeNotFoundError(name)
		}
		panic(err)
	}
	return nil
}
