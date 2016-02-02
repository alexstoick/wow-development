FROM golang:1.5.1

# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
Add . /go/src/github.com/alexstoick/wow/

# Get godeps from main repo
# RUN go get github.com/tools/godep

# Restore godep dependencies
# RUN godep restore

# Install
RUN go install github.com/alexstoick/wow/

# Setting up environment variables
ENV ENV dev

# My web app is running on port 8080 so exposed that port for the world
# EXPOSE 3000
# ENTRYPOINT ["/go/bin/budgie-backend"]
