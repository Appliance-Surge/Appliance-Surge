FROM golang:alpine

WORKDIR /app

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Run CompileDaemon to rebuild and restart on changes
CMD CompileDaemon -log-prefix=false -build="go build -o main ./cmd/appliance-surge/main.go" -command="./main"
