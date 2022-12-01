FROM golang:alpine



RUN go mod download

COPY app/*.go ./

RUN go get github.com/cosmtrek/air
