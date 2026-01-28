FROM golang:1.23-alpine

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/air-verse/air@v1.61.7

# Copy go mod files first for caching
COPY go.mod ./
COPY go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Download dependencies
RUN go mod tidy

EXPOSE 8080

# Use Air for hot reloading in development
CMD ["air", "-c", ".air.toml"]
