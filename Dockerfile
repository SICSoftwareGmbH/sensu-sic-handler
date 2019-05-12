# Build image
FROM golang:1.12-alpine

# Maintainer
MAINTAINER Florian Schwab <florian.schwab@sic.software>

# Upgrade system
RUN apk --no-cache --no-progress --update upgrade

# Install os dependencies
RUN apk --no-cache --no-progress --update add bash build-base curl git ca-certificates

# Install dep
RUN curl -sfL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install golangci-lint
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b /usr/local/bin v1.16.0

# Install goreleaser
RUN wget -q -O /tmp/goreleaser.tar.gz \
  https://github.com/goreleaser/goreleaser/releases/download/v0.106.0/goreleaser_Linux_arm64.tar.gz && \
  tar -xf /tmp/goreleaser.tar.gz -C /usr/local/bin && rm -rf /tmp/*

# Set the working directory
WORKDIR /go/src/sensu-sic-handler

# Default command
CMD ["bash"]
