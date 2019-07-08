# vessel-service/Dockerfile
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-service-vessel

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o shippy-service-vessel -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-vessel .

CMD ["./shippy-service-vessel"]
