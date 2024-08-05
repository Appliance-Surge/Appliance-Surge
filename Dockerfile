FROM golang:alpine

WORKDIR /app

# Install CompileDaemon
RUN apk add --no-cache postgresql-client nodejs npm entr \
    && go install github.com/githubnemo/CompileDaemon@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Copy the build-and-run script
COPY build-and-run.sh .

RUN chmod +x build-and-run.sh

# Run CompileDaemon to rebuild and restart on changes
CMD sh -c 'CompileDaemon -include="*.html" -log-prefix=false -command="./build-and-run.sh" -directory="./"'
