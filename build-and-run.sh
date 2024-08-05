#!/bin/sh

# Kill the previous instance if it's running
if pgrep main > /dev/null
then
    pkill main
fi

# Run npm build
npm run build

# Build and run the Go application
go build -o main ./cmd/appliance-surge/main.go

./main &
