FROM golang:1.23.2-alpine3.20 as builder
WORKDIR /app
RUN apk update
RUN apk upgrade
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o loadtester main.go

FROM scratch
EXPOSE 8080 8080
COPY --from=builder /app/loadtester .
CMD [ "./loadtester" ]


# FROM golang:1.23.2-alpine3.20
# # WORKDIR /app
# COPY . .
# RUN go build -o loadtester .
# CMD ["./loadtester"]
