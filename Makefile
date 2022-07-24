migrateup:
	migrate -path ./migrations -database "postgresql://default:secret@localhost:5432/go-crud1?sslmode=disable" -verbose up

migratedown:
	migrate -path ./migrations -database "postgresql://default:secret@localhost:5432/go-crud1?sslmode=disable" -verbose down