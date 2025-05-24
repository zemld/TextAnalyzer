FROM golang:alpine
WORKDIR /file-analyzer
COPY . .
RUN go mod tidy
CMD ["go", "run", "main/main.go"]