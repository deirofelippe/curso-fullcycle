FROM golang:1.16.0-stretch

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

COPY ./ ./

RUN go mod tidy

CMD ["go", "run", "main.go"]