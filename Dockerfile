FROM golang

# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
ADD . /go/src/github.com/alexstoick/wow/

RUN go get github.com/jinzhu/gorm
RUN go get github.com/robfig/cron
RUN	go get github.com/empatica/csvparser
RUN go get github.com/oleiade/reflections

# RUN go install github.com/alexstoick/wow/datafetch
RUN cd datafetch && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datafetch
# RUN go install github.com/alexstoick/wow/web
