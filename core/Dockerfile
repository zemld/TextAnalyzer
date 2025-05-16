FROM golang:alpine
WORKDIR /core
COPY . .
RUN go mod tidy
CMD ["go", "run", "main/main.go"]