start-dev-server:
	docker build -t pdf-print-server . -f Dockerfile.dev
	docker run --rm -p 4000:80 -v $$(pwd):/go/src/github.com/karnott/pdf-print-server pdf-print-server go run main.go
