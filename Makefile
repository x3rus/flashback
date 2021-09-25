

build: 
	go build -o flashback src/*.go 

test:
	go test -v src/*.go
