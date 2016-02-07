FROM golang

# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
ADD . /go/src/github.com/alexstoick/wow/

RUN go get github.com/jinzhu/gorm

RUN go install github.com/alexstoick/wow/datafetch
RUN go install github.com/alexstoick/wow/web
# Setting up environment variables
ENV ENV dev
