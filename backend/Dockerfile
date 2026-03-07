# Start from the official Go image
FROM golang:1.19-alpine

# Set the working directory
WORKDIR /app

# Install ent
RUN go install ent.io/ent/cmd/ent@latest

# Copy go mod file
COPY go.mod ./

# Copy source code
COPY . .

# Generate ent code
RUN /root/go/bin/ent generate ./ent/schema

# Download dependencies and build
RUN go mod tidy
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./main"]
