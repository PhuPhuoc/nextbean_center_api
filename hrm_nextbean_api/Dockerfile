FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .
RUN go build -o /hrm_nextbean_api

EXPOSE 8080

CMD ["/hrm_nextbean_api"]