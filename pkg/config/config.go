package config

type Config struct {
	// Time to wait inbetween each state check
	SleepTime string

	// Name of resource to query
	Resource string

	// Location of the resource.  Using a Kubernetes term, but this could be
	// translated as location
	Namespace string
}
