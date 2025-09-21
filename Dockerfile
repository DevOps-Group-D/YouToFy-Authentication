FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o YouToFy-Authentication .

FROM alpine:latest

WORKDIR /root/

# Copy only the compiled executable from the builder stage
COPY --from=builder /app/YouToFy-Authentication .

# Define the command to run the executable
CMD ["./YouToFy-Authentication"]
