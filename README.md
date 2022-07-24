# Go-Postgres

This project is simple CRUD application built in golang and using PostgreSQL as DB.

## Pre-requisite
1. Golang
2. API Tools like Postman / Insomnia / https://hoppscotch.io (need install extension proxy from it)
3. Install Latest stable docker
4. Running Docker
5. Run this command
```bash
$ docker compose up -d
$ make migrateup
$ go run main.go
```


## Cleanup
1. Run this command
```bash
$ docker compose down
$ docker image rm go-postgres_postgres
```