FROM golang

RUN mkdir -p /api

WORKDIR /api

COPY . .

RUN go build -o app main.go

ENTRYPOINT ["./app"]

EXPOSE 8080