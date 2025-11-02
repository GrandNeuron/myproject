# === Makefile для миграций и запуска ===
# строка подключения к базе
DB_DSN := "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"
# путь к утилите migrate
MIGRATE := migrate -path . -database $(DB_DSN)

.PHONY: migrate-new migrate-up migrate-down run

# команда для создания новой миграции
# пример вызова: make migrate-new NAME=init_calculations
migrate-new:
	@test -n "$(NAME)" || (echo "❌ Укажи имя: make migrate-new NAME=что_делаем"; exit 1)
	$(MIGRATE) create -ext sql -dir . $(NAME)

# применить все миграции
migrate-up:
	$(MIGRATE) up

# откатить последнюю миграцию
migrate-down:
	$(MIGRATE) down 1

# запустить сервер (с путём к main.go)
run:
	go run ../../cmd/main.go