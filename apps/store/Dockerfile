FROM golang:1.22

RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download
RUN mkdir -p cmd
COPY cmd ./cmd
RUN mkdir -p cmd
COPY internal ./internal
WORKDIR /opt/app/cmd/web

RUN CGO_ENABLED=0 GOOS=linux go build -o ../../store-server
WORKDIR /opt/app/
ENTRYPOINT ["./store-server"]