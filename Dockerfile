# Get Golang
FROM golang:1.20.4-alpine

# Version
LABEL version="1.1" name="Go API Example"

# Set the timezone (#RUN echo "UTC" > /etc/timezone)
RUN apk update && \
    apk add -U tzdata build-base && \
    cp /usr/share/zoneinfo/EST5EDT /etc/localtime && \
    echo "UTC" > /etc/timezone

# Set the working directory
WORKDIR /go/src/github.com/mrz1836/go-api

# Expose the port to the server
EXPOSE $API_SERVER_PORT

# Move the current code into the directory
COPY . /go/src/github.com/mrz1836/go-api

# Compile and build / Move the go application to the right path (/bin/) (hack)
RUN go build -i cmd/service/main.go && \
    go build cmd/service/main.go && \
    mv main /go/bin/

# Run the application
CMD ["main"]