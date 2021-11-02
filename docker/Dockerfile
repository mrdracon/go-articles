FROM golang:1.16.7

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the application
RUN go build -o go-articles cmd/main.go

RUN ls -la
RUN ls -la /app
EXPOSE 8080
CMD /app/go-articles

