CMD=go
BIN_PATH=bin
SRC_PATH=src

all: clean build install

build: 
	$(CMD) build -o $(BIN_PATH)/docker $(SRC_PATH)/*

install:
	cp bin/docker /usr/bin/docker
	cp bin/docker /usr/local/bin/docker
	mkdir -p /var/lib/docker/images
	mkdir -p /var/lib/docker/volumes
	mkdir -p /var/lib/docker/containers
	cp -r base /var/lib/docker/images/

uninstall:
	rm -rf /usr/bin/docker /usr/local/bin/docker
	rm -rf bin/docker

clean: uninstall
