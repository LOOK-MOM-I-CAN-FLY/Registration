# syntax=docker/dockerfile:1

# Этап сборки
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Копируем файлы модуля (go.mod и go.sum)
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники backend
COPY backend/ ./backend/

# Переходим в папку с файлом main.go
WORKDIR /app/backend/cmd

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Финальный образ на основе Alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированный бинарник
COPY --from=builder /app/backend/cmd/server .

# Копируем папку frontend для раздачи статики и шаблонов
COPY frontend/ ./frontend/

EXPOSE 8080

CMD ["./server"]
