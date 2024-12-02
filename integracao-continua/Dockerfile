FROM golang:1.23.3-bookworm

WORKDIR /app

RUN go mod init teste

COPY . .

RUN go build -o math

CMD [ "./math" ]