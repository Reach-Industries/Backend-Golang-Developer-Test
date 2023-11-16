FROM golang:1.18

WORKDIR /

COPY . .
RUN go build -o switcher .


FROM ubuntu:20.04

WORKDIR /root/

COPY --from=0 /switcher ./


CMD ["./switcher"]

