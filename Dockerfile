# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o /transaction-service

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /transaction-service /transaction-service
COPY .env /

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/transaction-service"]