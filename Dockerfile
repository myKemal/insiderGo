FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o insider-go .
EXPOSE 8090
CMD ["./insider-go"]
