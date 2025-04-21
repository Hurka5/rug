

all: clean build


build:
	go build -o rug ./main.go

clean:
	rm -f rug

test:
	go test -v ./... -timeout 3000ms
