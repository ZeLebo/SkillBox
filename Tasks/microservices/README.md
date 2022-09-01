To start the services separately need to run commands from cmd/proxy folder

Before starting on the local machine need to up the database

``docker run --name=box -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres``

``go run first.go`` - to run the first service

``go run second.go`` - to run the second service

``go run proxy.go`` - to run the proxy service

To run the services automatically need to run commands from cmd/app folder

change migration files

``migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up``