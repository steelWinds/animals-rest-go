## Build
FROM golang:alpine AS build

WORKDIR /home/app

ENV CGO_ENABLED 0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app

## Run
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]