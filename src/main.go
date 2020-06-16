package main

import (
	sv1b1 "test.go/src/server/v1beta1"
	sv1 "test.go/src/server/v1ga"
	scv1 "test.go/src/service/v1ga"
)

func main() {
	sv1b1.Print()
	sv1.Print()
	scv1.Print()
}
