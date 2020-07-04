# start from the golang 1.12 base image
FROM golang:1.12

# add maintainer info
LABEL maintainer='Rahul Verma <rv404674@gmail.com>'

# Set the Current Working Directory inside the container
WORKDIR /app

# Build Args
ARG LOG_DIR=/app/logs

# Create Log Directory.
# we will basically mount this directory (inside the docker container) to the one
# in our host machine.
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 


# install dependencies
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# build
RUN go build -o main .

#RUN chmod 777 scripts/build.sh && scripts/build.sh

EXPOSE 8080

# Declare volumes to mount
VOLUME [${LOG_DIR}]

CMD [ "./main" ]

