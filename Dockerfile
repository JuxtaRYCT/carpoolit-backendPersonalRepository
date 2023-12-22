FROM golang:latest as builder

WORKDIR /app

# Copying dependency files to the WORKDIR
COPY go.mod go.sum ./

# Downloading dependencies
RUN go mod download

# Copying source files to the WORKDIR
COPY . .

# Building the application
RUN CGO_ENABLED=0 GOOS=linux go build -o carpoolit-api

FROM alpine:latest

WORKDIR /root/

# Copying the binary from the builder container to the current directory
COPY --from=builder /app/carpoolit-api .
COPY --from=builder /app/.env .

# Exposing port 3000
EXPOSE 3000

# Running the application
CMD ["./carpoolit-api"]
