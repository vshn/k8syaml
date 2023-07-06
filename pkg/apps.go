package pkg

import apps "k8s.io/api/apps/v1"

func Deployment(name string) apps.Deployment {
	deployment := apps.Deployment{}
	deployment.Kind = "Deployment"
	deployment.APIVersion = "v1"
	deployment.Name = name
	return deployment
}

func StatefulSet(name string) apps.StatefulSet {
	statefulSet := apps.StatefulSet{}
	statefulSet.Kind = "StatefulSet"
	statefulSet.APIVersion = "v1"
	statefulSet.Name = name
	return statefulSet
}
