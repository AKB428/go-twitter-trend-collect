FROM golang:1.18 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags '-s -w' -o treco

# local only
RUN rm .env && mv .env_docker_local .env

# or scratch
FROM alpine
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/treco ./
COPY --from=builder /usr/src/app/.env .env
