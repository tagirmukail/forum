FROM golang:1.17.2-alpine AS builder

RUN apk update
RUN apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-w -s -X main.Version=$(git describe --tags --match 'v*' --always --abbrev=0)" \
     -o /app/forum cmd/forum/*.go

FROM scratch

COPY --from=builder /app/forum /app/forum

WORKDIR /app

EXPOSE 8080

CMD ["./forum"]