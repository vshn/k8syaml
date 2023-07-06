package pkg

import (
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Service(name string) core.Service {
	service := core.Service{}
	service.Kind = "Service"
	service.APIVersion = "v1"
	service.Name = name
	return service
}

func ResourceRequirements(cpuRequests string, cpuLimits string, memoryRequests string, memoryLimits string) core.ResourceRequirements {
	return core.ResourceRequirements{
		Requests: core.ResourceList{"cpu": resource.MustParse(cpuRequests), "memory": resource.MustParse(memoryRequests)},
		Limits:   core.ResourceList{"cpu": resource.MustParse(cpuLimits), "memory": resource.MustParse(memoryLimits)},
	}
}

func ProbeTCP(port int32, timeoutSeconds int32, periodSeconds int32, successThreshold int32, failureThreshold int32) *core.Probe {
	return &core.Probe{
		ProbeHandler: core.ProbeHandler{
			TCPSocket: &core.TCPSocketAction{
				Port: intstr.IntOrString{IntVal: port},
			},
		},
		TimeoutSeconds:   timeoutSeconds,
		PeriodSeconds:    periodSeconds,
		SuccessThreshold: successThreshold,
		FailureThreshold: failureThreshold,
	}
}
