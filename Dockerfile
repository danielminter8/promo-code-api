## We specify the base image we need for our
## go application
FROM golang:1.16
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .

RUN go mod download

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]



# FROM golang:1.15
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go mod download
# RUN go get github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="go build -o main ." --command="/app/main"