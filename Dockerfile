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

# install dependencies
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
# RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

####### Start a new image from scratch ###
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# COPY the Prebuilt binary file from the previous stage
COPY --from=builder /app/ .
COPY --from=builder /app/main/ .

# Expose port 8080 to the outside world
EXPOSE 8080

CMD [ "./main" ]