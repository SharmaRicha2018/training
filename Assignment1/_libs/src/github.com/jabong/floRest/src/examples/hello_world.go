package examples

import (
	"github.com/jabong/floRest/src/common/config"
	"github.com/jabong/floRest/src/common/constants"
	workflow "github.com/jabong/floRest/src/common/orchestrator"
)

type HelloWorld struct {
	id string
}

func (n *HelloWorld) SetID(id string) {
	n.id = id
}

func (n HelloWorld) GetID() (id string, err error) {
	return n.id, nil
}

func (a HelloWorld) Name() string {
	return "HelloWord"
}

func (a HelloWorld) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	//Business Logic
	appConfig, _ := config.GlobalAppConfig.ApplicationConfig.(*config.AppConfig)
	io.IOData.Set(constants.RESPONSE_HEADERS_CONFIG, appConfig.ResponseHeaders)
	return io, nil
}
