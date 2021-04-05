FROM golang:alpine3.13



RUN apk add --no-cache \
  bash \
  curl \
  && \
  rm -rf /var/cache/apk/*

ADD main.go /
ADD cdr-viewer.go /
ADD access_log /
ADD server.go /

RUN mkdir -p /var/log/asterisk/cdr-csv/
ADD Master.csv /var/log/asterisk/cdr-csv/Master.csv

CMD ["go", "run", "/server.go"]

