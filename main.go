package main

import (
	"fmt"
	k8syaml "github.com/vshn/k8syaml/pkg"
)

func main() {
	service := k8syaml.Service("myservice")

	err := k8syaml.WriteManifest(&service, "manifests")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
