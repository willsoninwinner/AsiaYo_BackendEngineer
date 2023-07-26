FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o AsiaYo_BackendEngineer

EXPOSE 8080

ENTRYPOINT ./AsiaYo_BackendEngineer