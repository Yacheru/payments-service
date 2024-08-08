docker-up:
	docker compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build

docker-down:
	docker compose down
	docker system prune -fa
