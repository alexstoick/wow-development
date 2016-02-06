FROM golang:1.5.1

# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
ADD . /go/src/github.com/alexstoick/wow/

# Get godeps from main repo
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore
RUN cd datafetch && godep restore

RUN go install github.com/alexstoick/wow/datafetch
RUN go install github.com/alexstoick/wow/web
# Setting up environment variables
ENV ENV dev
