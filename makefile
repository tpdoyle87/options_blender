up:
	docker compose up -d options_db
down:
	docker compose down options_db
build:
	docker compose build
running:
	docker ps -a
app-up:
	docker compose up options-app
app-down:
	docker compose down options-app
stop:
	docker stop
delete:
	docker container prune -f