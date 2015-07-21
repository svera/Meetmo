FROM golang:1.4

RUN go get \
	github.com/skelterjohn/rerun

RUN mkdir -p /go/src/github.com/svera/meetmo
ADD . /go/src/github.com/svera/meetmo

RUN ln -s /go/src/github.com/svera/meetmo /code

RUN cd /go/src/github.com/svera/meetmo && ls -l
RUN cd /go/src/github.com/svera/meetmo && go get && go build .

ENV PORT 8080
EXPOSE 8080

VOLUME ["/code"]
WORKDIR /code

CMD ["rerun", "--build", "--race", "github.com/svera/meetmo"]
