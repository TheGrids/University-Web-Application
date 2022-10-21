FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build main.go
CMD ["/app/main"]
LABEL Name=serv Version=0.0.1