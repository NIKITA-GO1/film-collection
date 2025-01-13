# Указываем базовый образ
FROM golang:1.17 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY cmd/go.mod cmd/go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY cmd/ ./

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o film-collection .

# Указываем образ запуска
FROM alpine:latest

# Копируем скомпилированное приложение из builder
COPY --from=builder /app/film-collection .

# Указываем команду для выполнения при запуске контейнера
CMD ["./film-collection"]