package resources

import (
	"github.com/automationbroker/bundle-controller/pkg/log"
)

type process interface {
	Process() []string
}

func (p podList) Process() []string {
	nameList := []string{}
	for _, pod := range p.Items {
		log.Debug(pod.Name)
		nameList = append(nameList, pod.Name)
	}
	return nameList
}
