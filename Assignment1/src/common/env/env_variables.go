package env

import (
	"os"
	"strings"
)

type OsEnviron struct {
	variables map[string]string
}

func (n *OsEnviron) initialise() {

	n.variables = make(map[string]string)
	envVars := os.Environ()
	for _, envVar := range envVars {
		if strings.Contains(envVar, "UPDATE_AUTH") {
			n.variables["UPDATE_AUTH"] = envVar[12:len(envVar)]
			continue
		}
		envVals := strings.SplitN(strings.Trim(envVar, " "), "=", 1)
		key := strings.Trim(envVals[0], " ")
		value := strings.Trim(envVals[1], " ")
		n.variables[key] = value
	}

}

func (n *OsEnviron) Get(key string) (value string, found bool) {
	value, found = n.variables[key]
	return value, found
}

func (n *OsEnviron) GetAll() map[string]string {
	return n.variables
}

//Global EnvVariable Singleton
var envVariables *OsEnviron = nil

func GetOsEnviron() *OsEnviron {
	if envVariables != nil {
		return envVariables
	}
	envVariables := new(OsEnviron)
	envVariables.initialise()
	return envVariables
}
