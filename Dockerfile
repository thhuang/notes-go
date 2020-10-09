FROM golang:1.13.5
LABEL maintainer="tzhsn.huang@gmail.com"

# Update apt-get
RUN apt-get update && apt-get clean

# Setup the project
WORKDIR /go/src
COPY get.sh /tmp/.
RUN bash /tmp/get.sh

# Entry
WORKDIR /go/src/github.com/Notes-Go
CMD /bin/bash
