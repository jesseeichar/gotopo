build:
	go build gotopo/geom
	
test:
	go test gotopo/geom/test
	
deps:
	go get github.com/stretchr/testify