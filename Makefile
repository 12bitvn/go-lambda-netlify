build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello main.go
deploy:
	netlify deploy --dir=site --functions=functions --prod