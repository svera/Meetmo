FROM golang:1.4

# Get rerun for live reloading of code inside the container
RUN go get \
	github.com/skelterjohn/rerun

# Create the full path for the project Go package
RUN mkdir -p /go/src/github.com/svera/meetmo
ADD . /go/src/github.com/svera/meetmo

# Symlink for mounting external code for live reloading
RUN ln -s /go/src/github.com/svera/meetmo /code
RUN cd /go/src/github.com/svera/meetmo && go get && go build .

# Setup the app TCP port
ENV PORT 8080
EXPOSE 8080

# Expose /code for mounting external code
VOLUME ["/code"]
WORKDIR /code

# Launch rerun, which will watch and recompile when source code changes
CMD ["rerun", "--build", "--race", "github.com/svera/meetmo"]
