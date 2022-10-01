package docker

import (
	"fmt"
	"os"
	"strings"

	"github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"
)

func (docker Docker) saveFileParameters(deployment *deployments.Deployment) error {
	dirPath := docker.generateDeploymentFileParameterPath(deployment.Name)
	fmt.Println("preparing: " + dirPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Errorf("error saving parameters to file: %f", err)
	}
	for _, parameter := range deployment.Parameters {
		path := docker.generateFileParameterPath(deployment.Name, parameter.Name)
		err := os.WriteFile(path, []byte(parameter.Value), os.ModePerm)
		fmt.Println("saving to " + path)
		if err != nil {
			fmt.Errorf("error saving parameters to file: %f", err)
		}
	}
	return nil
}

func (docker Docker) removeFileParameters(deploymentName string) {
	os.RemoveAll(docker.generateDeploymentFileParameterPath(deploymentName))
}

func (docker Docker) generateDeploymentFileParameterPath(deploymentName string) string {
	sb := strings.Builder{}
	sb.WriteString(strings.TrimRight(docker.bindLocation, "/"))
	sb.WriteString("/")
	sb.WriteString(deploymentName)
	return sb.String()
}

func (docker Docker) generateFileParameterPath(deploymentName string, parameterName string) string {
	sb := strings.Builder{}
	sb.WriteString(docker.generateDeploymentFileParameterPath(deploymentName))
	sb.WriteString("/")
	sb.WriteString(parameterName)
	sb.WriteString(".param")
	return sb.String()
}

func (docker Docker) generateFileParameterBindings(deploymentName string, parameters []*deployments.ParameterUsage) []string {
	binds := make([]string, 0)
	for _, parameter := range parameters {
		bindSource := docker.generateFileParameterPath(deploymentName, *parameter.Name) + ":"
		for _, bindTarget := range parameter.Files {
			bindString := bindSource + strings.TrimRight(bindTarget, "/") + ":ro"
			binds = append(binds, bindString)
			fmt.Println("binding via: " + bindString)
		}
	}
	return binds
}
