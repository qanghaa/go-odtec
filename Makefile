COMPOSE_FILE := ./deployments/proto.docker-compose.yml
gen-proto: 
	docker compose -f ${COMPOSE_FILE} up generate_pb_go --build