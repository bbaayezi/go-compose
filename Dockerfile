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

# Build the project (production)
RUN go build -o /go/bin/go-compose

# Testing
# RUN go run main.go

# Set program entry (production)
ENTRYPOINT [ "/go/bin/go-compose" ]