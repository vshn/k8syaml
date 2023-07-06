package main

import (
	"github.com/vshn/k8syaml/pkg"
)

func main() {
	service := pkg.Service("myservice")
	pkg.Write(&service)
	service = pkg.Service("")
	pkg.Write(&service)
	service = pkg.Service("")
	pkg.Write(&service)
}
