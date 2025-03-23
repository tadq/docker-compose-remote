FROM golang:1.23-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/service .

# Start fresh from a smaller image
FROM alpine:3.18
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/.env-prod /app/.env-prod
#RUN chmod +x /app/.env
#RUN /app/.env
#ENV $(cat /app/.env)
#RUN export $(cat /app/.env | xargs)

COPY --from=build_base /tmp/app/out/service /app/service

# This container exposes port to the outside world
EXPOSE 8008

# Set the Current Working Directory inside the container
WORKDIR /app

# Run the binary program produced by `go install`
CMD export $(cat /app/.env-prod) && /app/service
