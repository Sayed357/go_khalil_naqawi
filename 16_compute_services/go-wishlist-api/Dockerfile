FROM golang:1.22.4 AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN go build -o main .


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/main .

EXPOSE 1323

CMD ["./main"]
