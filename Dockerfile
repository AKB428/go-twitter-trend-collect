FROM golang:1.18 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -trimpath -ldflags '-s -w' -o treco

# local only
RUN rm .env && mv .env_docker_local .env
