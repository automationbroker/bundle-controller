package controller

import (
	"os"
	"reflect"
	"time"

	//      automationbrokerv1 "github.com/automationbroker/broker-client-go/client/clientset/versioned/typed/automationbroker.io/v1"
	// v1 "github.com/automationbroker/broker-client-go/pkg/apis/automationbroker.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
		SleepTime:   os.Getenv("SLEEPTIME"),
		Resource:    os.Getenv("RESOURCE"),
		Namespace:   os.Getenv("NAMESPACE"),
		BundleID:    os.Getenv("BUNDLEID"),
		BundleName:  os.Getenv("BUNDLENAME"),
		BundleParam: os.Getenv("BUNDLEPARAM"),
	}

	err = resources.NewCRDClient()
	if err != nil {
		log.Fatal(err)
	}

	if conf.BundleID != "" && conf.BundleName == "" {
		_, err = resources.Bundle.Bundles(conf.Namespace).Get(conf.BundleID, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}
	} else if conf.BundleName != "" {
		bundles, err := resources.Bundle.Bundles(conf.Namespace).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		log.Debugf("Looking for Bundle with Name '%v'", conf.BundleName)
		var count int
		for c, b := range bundles.Items {
			log.Debug(b.Spec.FQName)
			if b.Spec.FQName == conf.BundleName {
				conf.BundleID = b.Name
				log.Debugf("Found matching bundle '%v'", conf.BundleName)
				break
			}
			count = c
		}
		if count == len(bundles.Items)-1 {
			log.Fatalf("Didn't find Bundle '%v'", conf.BundleName)
		}
	} else {
		log.Fatal("BUNDLEID or BUNDLENAME is a required environment var")
	}

	return Controller{config: conf}
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

		// Check if the state has changed
		if !reflect.DeepEqual(oldState, currentState) {
			log.Info("** State Change **")

			log.Infof("Loading Bundle CRD: '%v'...", c.config.BundleID)
			b, err := resources.Bundle.Bundles(c.config.Namespace).Get(c.config.BundleID, metav1.GetOptions{})
			if err != nil {
				log.Errorf("Failed to load Bundle '%v' in namespace '%v'", c.config.BundleID, c.config.Namespace)
				log.Error(err)
			}
			status.UpdateState(currentState, b, c.config.Namespace, c.config.BundleParam)
		}

		log.Infof("Current list of items: %v", currentState)

		time.Sleep(sleep)
		oldState = currentState
	} // Controller Loop
}
