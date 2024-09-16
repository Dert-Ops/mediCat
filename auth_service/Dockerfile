FROM golang:1.20-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main ./cmd

FROM scratch

WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8080

CMD ["./main"]
