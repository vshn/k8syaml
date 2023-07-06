package pkg

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	"os"
	"reflect"
	"strings"
)

var (
	noname = 0
)

func Write(obj runtime.Object) {
	os.Mkdir("manifests", os.ModePerm)

	name := ""
	metadataValue := reflect.ValueOf(obj).Elem().FieldByName("ObjectMeta")
	if metadataValue.IsValid() {
		nameValue := metadataValue.FieldByName("Name")
		if nameValue.IsValid() {
			name = nameValue.String()
		}
	}
	if name == "" {
		name = fmt.Sprintf("noname%d", noname)
		noname = noname + 1
	}

	f, err := os.Create("manifests/" + strings.ToLower(obj.GetObjectKind().GroupVersionKind().Kind) + "-" + name + ".yaml")
	if err != nil {
		panic(err)
	}
	yp := printers.YAMLPrinter{}
	err = yp.PrintObj(obj, f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
