.PHONY: up logs down down_v

#starts db
up:
	docker-compose up -d

#shows db logs
logs:
	docker logs postgres

#stops db container
down:
	docker-compose down

#stops db container and deletes volumes
down_v:
	docker-compose down -v
