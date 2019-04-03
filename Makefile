build:
	go build -o gohello

clean:
	go clean

docker:
	docker build -t gohello .

.PHONY: build docker
	
