FROM golang:1.13.5
LABEL maintainer="thhuang@appcharmer.io"

# Update apt-get
RUN apt-get update && apt-get clean

# Setup the project
WORKDIR /go/src
COPY get.sh /tmp/.
RUN bash /tmp/get.sh

# Entry
WORKDIR /go/src/github.com/user
CMD /bin/bash