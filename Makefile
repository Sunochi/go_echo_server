# Makefile内でのみ適応される環境変数
export AWS_ACCESS_KEY_ID=minio_user
export AWS_SECRET_ACCESS_KEY=minio_password

init: docker-compose.up db.create db.schema.load

up:
	docker-compose exec app sh -c "go run main.go"

restart:
	docker-compose restart app
	docker-compose exec app sh -c "go run main.go"

log-check:
	docker-compose logs app

docker-compose.up:
	docker-compose up -d

format:
	go fmt ./...
	
lint:
	docker-compose exec app golangci-lint run ./...

clean.local:
	rm -f log/

package:
	docker-compose exec app sh -c "go build"
	
bash:
	docker-compose exec app sh

down.all:
	if [ -n "`docker ps -q`" ]; then docker kill `docker ps -q`; fi
	docker container prune -f

db.create:
	docker-compose exec db bash -c "mysql -uroot -hdb -e 'create database if not exists test;'"

db.schema.load:
	docker-compose exec db bash -c "mysql -uroot -hdb test < priv/db/structures/test.sql"