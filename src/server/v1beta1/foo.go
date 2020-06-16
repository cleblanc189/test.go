package v1beta1

import (
	"fmt"

	svc "test.go/src/service/v1beta1"
)

func Print() {
	fmt.Println("Printing fromi V1BETA1 subdir!")
	svc.Print()
}
