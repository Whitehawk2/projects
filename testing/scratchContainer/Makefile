.PHONY: install build docker clean

# Build target (statically linked)
build:
	CGO_ENABLED=0 go build -ldflags '-s -w' -o simplerouter .

# Install target (simply calls build)
install: build

# Clean target (removes the compiled binary)
clean:
	rm -f simplerouter

docker:
	install 
	docker build -t simplerouter:demo . -f Dockerfile-singleStage

# Default target (runs build)
all: build
