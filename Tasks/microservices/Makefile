postgres:
	docker run --name=box -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

postgres-stop:
	docker stop box

createdb:
	docker exec -it box createdb -U postgres -h localhost -p 5432 -O postgres test

dropdb:
	docker exec -it box dropdb -U postgres -h localhost -p 5432 test

migrateup:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migratedown:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

test:
	echo "testing make file"