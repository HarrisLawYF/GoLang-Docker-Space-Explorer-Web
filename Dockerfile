FROM golang:1.13

COPY . /go/src/SpaceApp
WORKDIR /go/src/SpaceApp/

# Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

#build the go app
RUN go build -o ./SpaceApp ./main.go

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["bee", "run"]