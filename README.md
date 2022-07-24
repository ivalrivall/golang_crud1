# Golang PostgreSQL

This project is simple CRUD application built in golang and using PostgreSQL as DB. Clone for use this project

## Pre-requisite
1. Go/Golang
2. API Tools like Postman / Insomnia / https://hoppscotch.io (need install extension proxy from it)
3. Install and running latest stable docker
4. Add database dependency by running this command
```bash
docker compose up -d
make migrateup
```
5. Running app
```bash
go run main.go
```


## Cleanup
1. Stop container and deleting
```bash
docker compose down
```
2. Delete image
```bash
docker image rm golang_crud1_postgres
```
