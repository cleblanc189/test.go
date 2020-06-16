module main

go 1.13

require (
	test.go/src/server/v1beta1 v0.0.0
	test.go/src/server/v1ga v0.0.0
	test.go/src/service/v1ga v0.0.0
)

replace (
	test.go/src/server/v1beta1 => ./server/v1beta1
	test.go/src/server/v1ga => ./server/v1ga
	test.go/src/service/v1ga => ./service/v1ga
)
