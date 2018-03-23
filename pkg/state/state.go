package state

import (
	"github.com/automationbroker/bundle-controller/pkg/log"
	resources "github.com/automationbroker/bundle-controller/pkg/resources"
)

// State - Represents the state of the resources we're monitoring
type State struct {
}

// NewState - Initialize controller state
func NewState() (*State, error) {
	return &State{}, nil
}

// CheckState - Check the state of resources
func (s State) CheckState(res string, namespace string) ([]string, error) {
	log.Info("CheckState - Look at the state of resources")

	resource, err := resources.Gather(res, namespace)
	if err != nil {
		log.Error("Failed to list pods in namespace '%v'", namespace)
		return nil, err
	}

	return resource.Process(), nil
}
