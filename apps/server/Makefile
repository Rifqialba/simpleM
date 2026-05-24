DB_URL=postgres://simplem:simplem_password@localhost:5433/simplem_db?sslmode=disable

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down

migrate-force:
	migrate -path database/migrations -database "$(DB_URL)" force 1