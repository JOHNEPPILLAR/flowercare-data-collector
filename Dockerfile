# syntax=docker/dockerfile:1

# Base image
FROM golang:bullseye as base
RUN apt-get update
#RUN export GO111MODULE=on
WORKDIR /app

# Build app
FROM base as build

# Install git
RUN apt-get install -y git

# Setup private repo
ARG GO_PRIVATE_REPO_KEY
RUN go env -w GOPRIVATE=github.com/JOHNEPPILLAR/*
RUN { echo "machine github.com"; echo "login JOHNEPPILLAR"; echo "password $GO_PRIVATE_REPO_KEY"; } >> "${HOME}"/.netrc \
    && chmod 600 $HOME/.netrc

# Copy go install files
COPY go.mod ./ 
COPY go.sum ./

# Download packages & build app
RUN go mod download
COPY *.go ./
RUN go build -o flowercare_collector

# Create final image
FROM base as final
RUN apt-get install -y tzdata sudo bluez bluetooth libbluetooth-dev libudev-dev
ENV TZ=Europe/London
COPY --from=build /app/flowercare_collector ${HOME}

# Setup start script
COPY docker-entrypoint.sh ./
CMD ["./docker-entrypoint.sh"]