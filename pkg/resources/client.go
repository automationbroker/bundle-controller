package resources

import (
	"github.com/automationbroker/bundle-controller/pkg/log"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *clientset.Clientset
var ClientConfig *rest.Config

// NewKubernetes - Initialize a cluster client
func NewKubernetes() error {
	clientConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Warningf("Failed to create a InternalClientSet: %v.", err)

		log.Info("Checking for a local Cluster Config")
		clientConfig, err = createClientConfigFromFile(homedir.HomeDir() + "/.kube/config")
		if err != nil {
			log.Error("Failed to create LocalClientSet")
			return err
		}
		log.Infof("Using local Cluster Config at '%v'.", homedir.HomeDir()+"/.kube/config")
	}

	clientset, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		log.Error("Failed to create LocalClientSet")
		return err
	}

	Client = clientset
	ClientConfig = clientConfig
	return err
}

func createClientConfigFromFile(configPath string) (*rest.Config, error) {
	clientConfig, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewDefaultClientConfig(*clientConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
