# syntax=docker/dockerfile:1
FROM golang:1.20

WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN go build  -o /app.o

EXPOSE 8080

# Run
CMD ["/app.o"]