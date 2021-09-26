

build: 
	go build -o flashback src/*.go 

test:
	go test -v src/*.go

lint:
	go vet  src/*.go
