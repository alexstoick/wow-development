FROM golang

# Setting up working directory
WORKDIR /go/src/github.com/alexstoick/wow/
ADD . /go/src/github.com/alexstoick/wow/

RUN go get github.com/jinzhu/gorm
RUN go get github.com/robfig/cron
RUN	go get github.com/empatica/csvparser
RUN go get github.com/oleiade/reflections
RUN go get github.com/gin-gonic/gin
RUN go get github.com/lib/pq

RUN cd datafetch && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datafetch
RUN cd web && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o web
