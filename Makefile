docker-up:
	docker compose up --build

docker-down:
	docker compose down

docker-down-v:
	docker compose down -v

docker-exec-db:
	docker exec -it groupie-tracker-db-1 psql -U divine -d groupie-tracker