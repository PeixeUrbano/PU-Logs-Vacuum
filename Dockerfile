# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8

# Directory inside the container to store all application and then make it the working directory.
RUN mkdir -p /go/src/logs-vacuum
WORKDIR /go/src/logs-vacuum

# Copy the current directory into the container.
COPY . /go/src/logs-vacuum

# Download and install go required third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install