package state

import (
	"reflect"

	v1 "github.com/automationbroker/broker-client-go/pkg/apis/automationbroker.io/v1"
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
		log.Errorf("Failed to gather resource '%v' in namespace '%v'", res, namespace)
		return nil, err
	}

	return resource.Process(), nil
}

var paramFound bool
var planIndex, paramIndex int

// UpdateState - Update the Bundle CRD
func (s State) UpdateState(currentState []string, bundle *v1.Bundle, namespace string, bundleParam string) error {
	if paramFound {
		// Check if Bundle CRD matches
		log.Debug("Checking if Bundle CRD parameters match.")
		if !reflect.DeepEqual(bundle.Spec.Plans[planIndex].Parameters[paramIndex].Enum, currentState) {
			// Update the Bundle Enum with the currentState
			updateBundle(currentState, bundle, namespace)
		}
		log.Info("Bundle parameters updated!")
	} else {
		for p, plan := range bundle.Spec.Plans {
			for a, param := range plan.Parameters {
				if param.Name == bundleParam {
					// Cache the location of the param we want to update
					planIndex = p
					paramIndex = a
					paramFound = true

					updateBundle(currentState, bundle, namespace)
				}
			}
		}
	}
	return nil
}

func updateBundle(currentState []string, bundle *v1.Bundle, namespace string) {
	log.Debug("Updating Bundle CRD")
	bundle.Spec.Plans[planIndex].Parameters[paramIndex].Enum = currentState
	_, err := resources.Bundle.Bundles(namespace).Update(bundle)
	if err != nil {
		// Don't fail on the update
		// TODO: try again later on a failure
		log.Errorf("Failed to update Bundle: %v", err)
	}
}
