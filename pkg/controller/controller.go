package controller

import (
	"reflect"
	"time"

	"github.com/automationbroker/bundle-controller/pkg/config"
	"github.com/automationbroker/bundle-controller/pkg/log"
	"github.com/automationbroker/bundle-controller/pkg/resources"
	"github.com/automationbroker/bundle-controller/pkg/state"
)

type Controller struct {
	config config.Config
}

func CreateController() Controller {
	log.NewLog()
	err := resources.NewKubernetes()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("===== Starting Bundle Controller =====")

	conf := config.Config{
		SleepTime: "1s",
		Resource:  "pod",
		Namespace: "ansible-service-broker",
	}

	c := Controller{config: conf}

	return c
}

func (c *Controller) Start() {
	status, err := state.NewState()
	if err != nil {
		log.Fatal(err)
	}
	sleep, err := time.ParseDuration(c.config.SleepTime)
	if err != nil {
		log.Fatal(err)
	}

	var currentState []string
	oldState, err := status.CheckState(c.config.Resource, c.config.Namespace)
	if err != nil {
		log.Fatal(err)
	}
	for {
		currentState, err = status.CheckState(c.config.Resource, c.config.Namespace)
		if err != nil {
			log.Fatal(err)
		}

		if !reflect.DeepEqual(oldState, currentState) {
			// Check if Bundle CRD matches
			log.Debug("Checking if Bundle CRD parameters match.")
		}

		log.Infof("Current list of items: %v", currentState)

		time.Sleep(sleep)
		oldState = currentState
	} // Controller Loop
}
