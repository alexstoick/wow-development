FROM golang

VOLUME testing-volume /output
# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
ADD . /go/src/github.com/alexstoick/wow/

RUN go get github.com/jinzhu/gorm

# RUN go install github.com/alexstoick/wow/datafetch
RUN cd datafetch && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /output/datafetch
# RUN go install github.com/alexstoick/wow/web
RUN echo "ls /output"
