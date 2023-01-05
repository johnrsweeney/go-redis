FROM golang:1.19

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go install .

ENTRYPOINT ["go-redis"]

