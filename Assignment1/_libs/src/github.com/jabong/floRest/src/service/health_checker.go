package service

import (
	"github.com/jabong/floRest/src/common/config"
)

type ServiceHealthCheck struct {
}

func (n ServiceHealthCheck) GetName() string {
	return "service"
}

func (n ServiceHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"version": config.GlobalAppConfig.AppVersion,
	}
}
