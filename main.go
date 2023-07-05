package main

import (
	"fmt"
	"github.com/vshn/k8syaml/pkg"
)

func main() {
	service := pkg.Service("myservice")

	err := pkg.WriteManifest(&service, "manifests")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
