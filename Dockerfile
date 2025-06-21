FROM golang:1.24
RUN mkdir -p /app
WORKDIR /app
COPY . /app
EXPOSE 8080
RUN go build -o app
CMD ["./app"]
