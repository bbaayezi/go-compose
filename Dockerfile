FROM golang

# Setup working dir
WORKDIR /go/src

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy remaining code
COPY . .

# Build the project
RUN go build -o /go/bin/go-compose

# Set program entry
ENTRYPOINT [ "/go/bin/go-compose" ]