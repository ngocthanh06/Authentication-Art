# Đặt tên cho các biến để sử dụng trong các lệnh
MIGRATE_COMMAND := migrate
MIGRATION_DIR := internal/database/migrations
MIGRATION_EXT := sql
URL_DATABASE := postgres://user:pass@127.0.0.1:5432/authentication?sslmode=disable
DOCKER_COMPOSE := docker-compose
BASH_COMMAND := $(DOCKER_COMPOSE) exec app bash
MAIN := app/authentication/main.go

# Tạo một mục tiêu để tạo một tệp migration mới
.PHONY: migrate-create

create-repository:
	@if [ -z ]


migrate-create:
	@if [ -z "$(name)" ]; then \
  		echo "Error: Please provide a migration name using 'make migrate-create name==<migration_name>'";\
  		exit 1;\
	fi
	$(MIGRATE_COMMAND) create -ext $(MIGRATION_EXT) -dir $(MIGRATION_DIR) -seq $(name)

# Thêm một mục tiêu để chạy migrations (tùy chọn)
migrate-up:
	$(MIGRATE_COMMAND) -database ${URL_DATABASE} -path $(MIGRATION_DIR) up

migrate-down:
	$(MIGRATE_COMMAND) -database ${URL_DATABASE} -path $(MIGRATION_DIR) down

down:
	$(DOCKER_COMPOSE) down

reset:
	$(DOCKER_COMPOSE) reset

build:
	$(DOCKER_COMPOSE) up -d --build

bash:
	$(DOCKER_COMPOSE) exec app bash

init-jwt-key:
	$(BASH_COMMAND) -c "go run $(MAIN) init-jwt-key"
