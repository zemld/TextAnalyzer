FROM golang:alpine
WORKDIR /file-storager
COPY . .
RUN go mod tidy
CMD ["go", "run", "main/main.go"]