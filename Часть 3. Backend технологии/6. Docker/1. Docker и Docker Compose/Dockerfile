FROM golang:1.24-bookworm

WORKDIR /app

COPY . .

RUN go build -o /app/exe main.go

CMD ["/app/exe" ]
