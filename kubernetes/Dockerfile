FROM golang:1.23.2-bookworm

RUN apt update \
    && useradd -u 1000 devfc

USER devfc

WORKDIR /home/devfc/app

COPY --chown=devfc:devfc ./ ./

RUN go build -o server .

CMD [ "./server" ]