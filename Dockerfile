# Build image
FROM golang:1.12-alpine

# Maintainer
MAINTAINER Florian Schwab <me@ydkn.de>

# Upgrade system
RUN apk --no-cache --no-progress --update upgrade

# Install dependencies
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

# Set the working directory
WORKDIR /go/src/sensu-sic-handler

# Copy in the application code
COPY ./ /go/src/sensu-sic-handler

# Install dependencies
RUN dep ensure

# Default command
CMD ["bash"]
