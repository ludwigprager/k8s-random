#FROM alpine:3.12
FROM golang:alpine3.13



RUN apk add --no-cache \
  bash \
  curl \
  file\
  iproute2 \
  ipset \
  iputils \
  libc6-compat \
  net-snmp-tools \
  netcat-openbsd \
  nftables \
  ngrep \
  nmap \
  nmap-nping \
  openssl \
  socat \
  strace \
  tcpdump \
  tcptraceroute \
  util-linux \
  vim \
  && \
  rm -rf /var/cache/apk/*

ADD main.go /
ADD cdr-viewer.go /
ADD access_log /
ADD liveness.go /

RUN mkdir -p /var/log/asterisk/cdr-csv/
ADD Master.csv /var/log/asterisk/cdr-csv/Master.csv

#CMD ["/bin/bash","-l"]
# go run cdr-viewer.go &
#CMD ["go", "run", "/cdr-viewer.go"]
CMD ["go", "run", "/liveness.go"]

