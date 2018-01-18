package main

import (
	"common/appconfig"
	"common/appconstant"
	"common/dao/mysqlconn"
	//"common/notification"
	"fmt"
	"github.com/jabong/floRest/src/service"
	"hello"
	//model "selfhelp/models"
)

//main is the entry point of the florest web service
func main() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)

	registerConfig()

	registerErrors()
	registerAllApis()

	mysqlconn.Initialize()
	//model.GetTemplateText()
	//model.GetSMSExtConfigSettings()
	webserver.Start()
}

func registerAllApis() {
	service.RegisterApi(new(hello.HelloApi))

}

func registerConfig() {
	service.RegisterConfig(new(appconfig.AppConfig))
}

func registerErrors() {
	service.RegisterHttpErrors(appconstant.AppErrorCodeToHttpCodeMap)
}

func overrideConfByEnvVariables() {
	service.RegisterGlobalEnvUpdateMap(appconfig.MapEnvVariables())
}
