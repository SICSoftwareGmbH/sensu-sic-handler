# Build image
FROM golang:1.12-alpine

# Maintainer
MAINTAINER Florian Schwab <florian.schwab@sic.software>

# Upgrade system
RUN apk --no-cache --no-progress --update upgrade

# Install os dependencies
RUN apk --no-cache --no-progress --update add bash build-base git ca-certificates

# Install dep
RUN go get github.com/golang/dep/cmd/dep

# Install cobra
RUN go get github.com/spf13/cobra/cobra

# Install golint
RUN go get golang.org/x/lint/golint

# Install errcheck
RUN go get github.com/kisielk/errcheck

# Install goconst
RUN go get github.com/jgautheron/goconst/cmd/goconst

# Install ineffassign
RUN go get github.com/gordonklaus/ineffassign

# Install goreleaser
RUN wget -q -O /tmp/goreleaser.tar.gz \
  https://github.com/goreleaser/goreleaser/releases/download/v0.106.0/goreleaser_Linux_arm64.tar.gz && \
  tar -xf /tmp/goreleaser.tar.gz -C /usr/local/bin && rm -rf /tmp/*

# Set the working directory
WORKDIR /go/src/sensu-sic-handler

# Default command
CMD ["bash"]
