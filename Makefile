.PHONY: proto build docker-build docker-run

# Генерация proto файлов
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/auth/auth_service.proto

# Сборка приложения
build:
	go build -o bin/auth-service cmd/auth-service/main.go

# Сборка Docker образа
docker-build:
	docker build -t auth-service:latest .

# Запуск в Docker
docker-run:
	docker run -p 50051:50051 auth-service:latest

# Очистка
clean:
	rm -rf bin/
	rm -rf pkg/auth/*.pb.go 