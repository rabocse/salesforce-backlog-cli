# syntax=docker/dockerfile:1

 ## While optional, the above directive instructs the Docker builder what syntax to use when parsing the Dockerfile, and allows older Docker versions with BuildKit enabled to upgrade the parser before starting the build.
 ## It is recommended using docker/dockerfile:1, which always points to the latest release of the version 1 syntax. BuildKit automatically checks for updates of the syntax before building, making sure you are using the most current version.


##  Base image we would like to use for our application.
FROM golang:1.16-alpine

LABEL maintainer="Alexander Escobar"
LABEL email="alexrabocse.me@gmail.com"

## Create a directory. This also instructs Docker to use this directory as the default destination for all subsequent commands.
WORKDIR /salesforce-backlog-cli 

RUN apk update
RUN apk add git
RUN apk upgrade

## Get necessary modules to compile the code.
COPY go.mod go.sum ./

## This works exactly the same as if we were running go locally on our machine, but this time these Go modules will be installed into a directory inside the image.
RUN go mod download

## Copy source code into the image.
COPY main.go ./
COPY sftool/ ./sftool/

## Compile the application.
RUN go build main.go

## Execute when is used to start a container
CMD ["sh"]