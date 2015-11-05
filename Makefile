build:
	go build gotopo/geom
	
test:
	go test gotopo/geom/test gotopo/geom/coords64
	
deps:
	go get github.com/stretchr/testify