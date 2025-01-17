# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bems-api

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bems-api /bems-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bems-api"]