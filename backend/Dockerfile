# Builder stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy all source files first
COPY . .

# Install git, which is needed for go mod tidy
RUN apk add --no-cache git

# Tidy downloads dependencies and creates a correct go.sum
RUN go mod tidy



# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

# Final stage
FROM golang:1.19-alpine

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /main /main

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["/main"]
