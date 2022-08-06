FROM golang:1.19 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags '-s -w' -o treco

# local only
RUN rm .env && mv .env_docker_local .env

FROM scratch
WORKDIR /usr/src/app
COPY --from=builder ["/usr/src/app/treco" ,"/usr/src/app/.env", "./"]
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt