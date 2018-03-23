package resources

import (
	"github.com/automationbroker/bundle-controller/pkg/log"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Gather - Gather resources based on config setting
func Gather(config string, namespace string) (process, error) {
	if config == "pod" {
		pods, err := getPodList(namespace)
		if err != nil {
			log.Errorf("Failed listing pods")
			return nil, err
		}
		return podList{pods}, nil
	}
	return nil, nil
}

type podList struct {
	*v1.PodList
}

func getPodList(namespace string) (*v1.PodList, error) {
	pod, err := Client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	return pod, err
}
