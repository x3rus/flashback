

build: 
	go build -o flashback src/*.go 

test:
	go test -v src/*.go

lint:
	go vet  src/*.go

localdev: build
	ALBUMDIRS="data/pic-sample/" ./flashback
