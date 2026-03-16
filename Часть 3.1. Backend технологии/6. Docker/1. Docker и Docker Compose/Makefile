include .env
export

service-run:
	go run main.go

service-deploy:
	docker compose up --build -d application

service-undeploy:
	docker compose down application

postgres-up:
	docker run -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p 5432:5432 -v out/pgdata -d postgres:18-bookworm
