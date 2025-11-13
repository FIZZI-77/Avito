
FROM golang:1.24-alpine AS build

WORKDIR /app
RUN apk add --no-cache bash git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./src/cmd/server \
    && chmod +x server


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

COPY .env ./

CMD ["./server"]
