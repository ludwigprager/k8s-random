#FROM alpine:3.12
FROM golang:alpine3.13



RUN apk add --no-cache \
  bash \
  bind-tools \
  bridge-utils \
  apache2-utils \
  conntrack-tools \
  curl \
  dhcping \
  ethtool \
  file\
  fping \
  iperf \
  iproute2 \
  ipset \
  iptables \ 
  iptraf-ng \
  iputils \
  jq \
  libc6-compat \
  liboping \
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

#CMD ["/bin/bash","-l"]
# go run cdr-viewer.go &
CMD ["go", "run", "/cdr-viewer.go"]

