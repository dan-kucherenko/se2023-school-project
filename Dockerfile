FROM golang:1.20-alpine

WORKDIR /ses-project-app

COPY . .

RUN go build -o ses-project-app

EXPOSE 8888

CMD ["./ses-project-app"]