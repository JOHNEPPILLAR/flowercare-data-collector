# syntax=docker/dockerfile:1

# Base image
FROM golang:alpine as base-build
WORKDIR /app

# Update and install git
RUN apk update && apk add git

# Setup private repo
ARG GO_PRIVATE_REPO_KEY
RUN go env -w GOPRIVATE=github.com/JOHNEPPILLAR/*
RUN { echo "machine github.com"; echo "login JOHNEPPILLAR"; echo "password $GO_PRIVATE_REPO_KEY"; } >> "${HOME}"/.netrc \
    && chmod 600 $HOME/.netrc

# Copy go install files
COPY go.mod ./ 
COPY go.sum ./

# Download packages
RUN go mod download && go mod verify

# Build app
COPY . .
RUN go build -o /flowercare cmd/main.go

# Start fresh from a smaller image
FROM alpine:latest

# Add certs and tzdata
RUN apk update && apk add ca-certificates && apk add tzdata bluez bluetooth libbluetooth-dev libudev-dev

# Set timezone and copy built app
ENV TZ=Europe/London
COPY --from=base-build /flowercare ./flowercare

# Setup start script
COPY docker-entrypoint.sh ./
CMD ["./docker-entrypoint.sh"]