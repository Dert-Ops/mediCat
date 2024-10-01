FROM golang:1.21.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./

FROM scratch
WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8082
CMD ["./main"]
