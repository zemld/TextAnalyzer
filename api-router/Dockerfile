FROM golang:alpine
WORKDIR /api
COPY . .
RUN go mod tidy
CMD ["go", "run", "main/main.go"]