FROM golang:1.23-alpine

WORKDIR /app

# Install build dependencies and Air
RUN apk add --no-cache gcc musl-dev git \
    && go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
