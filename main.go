package main

import (
	"log"
	"os"

	"github.com/keptn/go-utils/pkg/sdk"
	"github.com/ortelius/keptn-ortelius-service/handler"
	"github.com/sirupsen/logrus"
)

const getSliTriggeredEvent = "sh.keptn.event.get-sli.triggered"
const actionTriggeredEvent = "sh.keptn.event.action.triggered"
const serviceName = "keptn-ortelius-service"
const envVarLogLevel = "LOG_LEVEL"

func main() {
	if os.Getenv(envVarLogLevel) != "" {
		logLevel, err := logrus.ParseLevel(os.Getenv(envVarLogLevel))
		if err != nil {
			logrus.WithError(err).Error("could not parse log level provided by 'LOG_LEVEL' env var")
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			logrus.SetLevel(logLevel)
		}
	}

	log.Printf("Starting %s", serviceName)

	log.Fatal(sdk.NewKeptn(
		serviceName,
		sdk.WithTaskHandler(
			actionTriggeredEvent,
			handler.NewActionTriggeredEventHandler()),
		sdk.WithTaskHandler(
			getSliTriggeredEvent,
			handler.NewGetSliEventHandler()),
		sdk.WithLogger(logrus.New()),
	).Start())
}
