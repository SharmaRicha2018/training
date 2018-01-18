package appconfig

import (
	"fmt"
	florest_config "github.com/jabong/floRest/src/common/config"
	"github.com/jabong/floRest/src/common/logger"
	"os"
)

type AppConfig struct {
	Hello           florest_config.Application
	PromotionConf   PromotionConf
	EncrpServConf   EncrpServConf
	AmqpUrl         string
	SmsTemplateText string
}

type PromotionConf struct {
	CreateVoucher string
}

type EncrpServConf struct {
	ReqTimeout      int
	EndpointEncrypt string
	EndpointDecrypt string
	Host            string
}

var Config *AppConfig

func LoadAppConfig() {
	c := florest_config.GlobalAppConfig.ApplicationConfig
	appConfig, ok := c.(*AppConfig)
	if !ok {
		msg := fmt.Sprintf("APP Config Not correct %+v", c)
		logger.Error(msg)
		//return nil, errors.New(msg)
	}
	Config = appConfig
	//return appConfig, nil
}

// MapEnvVariables -> Returns map of config values to be replaced by environment variables
func MapEnvVariables() map[string]string {
	overrideVar := make(map[string]string)
	overrideVar["ApplicationConfig.PromotionConf.CreateVoucher"] = "CREATE_VOUCHER_URL"
	overrideVar["ApplicationConfig.EncrpServConf.EndpointEncrypt"] = "ENCRYPT_ENDPOINT"
	overrideVar["ApplicationConfig.EncrpServConf.EndpointDecrypt"] = "DECRYPT_ENDPOINT"
	overrideVar["ApplicationConfig.EncrpServConf.Host"] = "ENCRYPTION_SERVICE_HOST"
	overrideVar["ApplicationConfig.AmqpUrl"] = "AMQP_URL"
	checkEnv(overrideVar)
	return overrideVar
}

// checkEnv -> Checks environment variable availability in map, deletes entry if doesn't exist.
func checkEnv(override map[string]string) {
	for key, value := range override {
		if os.Getenv(value) == "" {
			delete(override, key)
		}
	}
}
