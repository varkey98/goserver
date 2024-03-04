# syntax=docker/dockerfile:1

FROM ubuntu
RUN mkdir -p /goagent-src
COPY goagent-src/ /goagent-src/
RUN ls -la /goagent-src/*


FROM golang:1.19

# Set destination for COPY
WORKDIR /app


# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY GoServer/*.go ./
COPY GoServer/libtraceable.so ./
COPY GoServer/config.yaml ./

# Build
RUN go build -tags 'traceable_filter' -o /app.o

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8087

# Run
CMD ["/app.o"]