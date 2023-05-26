FROM golang:1.20-alpine

WORKDIR /se-project-app

COPY . .

RUN go build -o se-project-app

EXPOSE 8888

CMD ["./se-project-app"]