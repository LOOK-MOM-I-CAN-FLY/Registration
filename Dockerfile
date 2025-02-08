FROM golang:1.21

WORKDIR /app

COPY backend backend
COPY frontend frontend
COPY go.mod go.sum ./

RUN cd backend && go build -o server cmd/main.go

CMD ["/app/backend/server"]
