FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# copy source files
COPY *.go ./
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /viter main.go

# Deploy the application binary into a lean image
FROM alpine:3.14

WORKDIR /

COPY --from=build-stage /viter /viter

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/viter"]
