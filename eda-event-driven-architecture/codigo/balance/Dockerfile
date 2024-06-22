FROM golang:1.20

WORKDIR /app

RUN apt update -y \
    && apt install -y librdkafka-dev git \
    # Install Delver (Go Debugger)
    && cd /home \
    && git clone https://github.com/go-delve/delve \
    && cd delve \
    && go install github.com/go-delve/delve/cmd/dlv@v1.22.1

CMD [ "sleep", "inf" ]