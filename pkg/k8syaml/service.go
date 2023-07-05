package k8syaml

import core "k8s.io/api/core/v1"

func Service(name string) core.Service {
	service := core.Service{}
	service.Kind = "Service"
	service.APIVersion = "v1"
	service.Name = name
	return service
}
