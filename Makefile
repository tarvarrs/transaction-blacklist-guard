.PHONY: up logs down down_v run

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

#runs backend
run:
	go run cmd/app/main.go
