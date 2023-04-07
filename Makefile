.PHONY: setup-db run-insert run

setup-db:
	docker stop mysql-local || true
	docker rm mysql-local || true
	docker run --name mysql-local -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=game_db -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 -d mysql:latest

run-insert:
	go run create_data/data_insert.go

run: 
	go run main.go

run-all: setup-db run-insert
