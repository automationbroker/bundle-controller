package resources

import (
	"github.com/automationbroker/bundle-controller/pkg/log"

	crdclientset "github.com/automationbroker/broker-client-go/client/clientset/versioned"
	automationbrokerv1 "github.com/automationbroker/broker-client-go/client/clientset/versioned/typed/automationbroker.io/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *clientset.Clientset
var ClientConfig *rest.Config
var Bundle automationbrokerv1.AutomationbrokerV1Interface

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

func NewCRDClient() error {
	// NOTE: Both the external and internal client object are using the same
	// clientset library. Internal clientset normally uses a different
	// library
	clientConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Warning("Failed to create a InternalClientSet: %v.", err)

		log.Debug("Checking for a local Cluster Config")
		clientConfig, err = createClientConfigFromFile(homedir.HomeDir() + "/.kube/config")
		if err != nil {
			log.Error("Failed to create LocalClientSet")
			return err
		}
	}

	crd, err := crdclientset.NewForConfig(clientConfig)
	if err != nil {
		log.Error("Failed to create LocalClientSet")
		return err
	}

	Bundle = crd.AutomationbrokerV1()
	return nil
}
