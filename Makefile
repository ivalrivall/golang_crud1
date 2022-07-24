migrateup:
	migrate -path ./migrations -database "postgresql://godbusr:secret@localhost:5432/gocrud1?sslmode=disable" -verbose up

migratedown:
	migrate -path ./migrations -database "postgresql://godbusr:secret@localhost:5432/gocrud1?sslmode=disable" -verbose down