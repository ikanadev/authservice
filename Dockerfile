FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download
COPY . /app
RUN /app/wait-for-it.sh -t 2 authservicedb:5432 -- echo "DB is ok"
RUN CGO_ENABLED=0 go build -o main /app/pkg/api/main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/templates/email.html ./templates/
EXPOSE 8000
CMD ["./main"]
