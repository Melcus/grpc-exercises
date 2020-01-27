composeFile = docker-compose.yml
composeBinaryFile = docker-compose.binary.yml

.PHONY: binary

binary:
	rm -f binary
	docker-compose -f $(composeBinaryFile) build
	docker-compose -f $(composeBinaryFile) up
	docker cp synapse_builder:/app/binary ./binary

build: # build containers
	docker-compose -f $(composeFile) build

start: #start previously stopped containers
	docker-compose -f $(composeFile) up -d --no-recreate

stop: # stop containers
	docker-compose -f $(composeFile) stop

clean: # delete container + data volumes
	docker-compose -f $(composeFile) rm -f