FROM golang:1.24

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -gcflags="all=-N -l" -o timekeeper ./cmd/main.go

EXPOSE 8080 40000

CMD ["dlv", "exec", "./timekeeper", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]
