# NOTE If we dont optimize our image it's size will be 1.5Gb.
# hence use multi-stage build
# With multistage builds you can use multiple FROM statements in your dockerfile. Each FROM instruction
# can use a different base, and each of them begins a new stage of build.

# The first stage will use golang:latest image and build our application. The second stage will use a very
# lightweight Alpine Linux Image and will only contain the binary executable built by first stage.

# start from the golang 1.12 base image
FROM golang:latest as builder

# add maintainer info
LABEL maintainer='Rahul Verma <rv404674@gmail.com>'

# Set the Current Working Directory inside the container
WORKDIR /app

# copy the source from the current directory to the working directory inside the container
# this will copy both go.mod go.sum as well. So no need to copy again
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
# RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

####### Start a new image from scratch ###
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# COPY the Prebuilt binary file from the previous stage
# "COPY --from=builder /app/main/ ." No need to do this as we are copy whole app including the executables.
# we need the whole app as well, because apart from executable our environment variables are stored in a
# .env file.
COPY --from=builder /app/ .

# Expose port 8080 to the outside world
# This port should be the same one, as exposed by the app server (goRubu).
EXPOSE 8080

CMD [ "./main" ]