package docker

import (
	"fmt"
	"os"
	"strings"

	"github.com/acul009/control-panel-2/src/api/deployments/gen/deployments"
)

func (docker Docker) saveFileParameter(deploymentName string, parameterName string, parameters []*deployments.Parameter) error {
	dirPath := docker.generateDeploymentFileParameterPath(deploymentName)
	fmt.Println("preparing: " + dirPath)
	err := os.MkdirAll(dirPath, os.ModePerm)

	if err != nil {
		return fmt.Errorf("error saving parameters to file: %f", err)
	}

	value, err := docker.getParameterValue(parameters, parameterName)

	if err != nil {
		return fmt.Errorf("error saving parameters to file: %f", err)
	}

	path := docker.generateFileParameterPath(deploymentName, parameterName)

	err = os.WriteFile(path, []byte(value), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error saving parameters to file: %f", err)
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

func (docker Docker) getParameterValue(parameters []*deployments.Parameter, parameterName string) (string, error) {
	for _, parameter := range parameters {
		if parameter.Name == parameterName {
			return parameter.Value, nil
		}
	}
	return "", fmt.Errorf("The Parameter %s is not defined", parameterName)
}

func (docker Docker) generateFileParameterBindings(deploymentName string, parameterBindings []*deployments.ParameterUsage, parameters []*deployments.Parameter) []string {
	binds := make([]string, 0)
	for _, parameterBinding := range parameterBindings {

		docker.saveFileParameter(deploymentName, *parameterBinding.Name, parameters)

		bindSource := docker.generateFileParameterPath(deploymentName, *parameterBinding.Name) + ":"
		for _, bindTarget := range parameterBinding.Files {
			bindString := bindSource + strings.TrimRight(bindTarget, "/") + ":ro"
			binds = append(binds, bindString)
		}
	}
	return binds
}

func (docker Docker) generateEnvParameterBindings(deploymentName string, parameterBindings []*deployments.ParameterUsage, parameters []*deployments.Parameter) ([]string, error) {
	binds := make([]string, 0)
	for _, parameterBinding := range parameterBindings {
		for _, bindTarget := range parameterBinding.Environment {
			value, err := docker.getParameterValue(parameters, *parameterBinding.Name)

			if err != nil {
				return []string{}, fmt.Errorf("Error generating parameterbinding for environment: %f", err)
			}

			bindString := bindTarget + "=" + value
			binds = append(binds, bindString)
		}
	}
	return binds, nil
}
