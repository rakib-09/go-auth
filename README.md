# go-auth

A simple web application using golang.

### Stack:
1. Echo Framework
2. Mysql
3. Docker

## Setup

1. Open Terminal and Run `cp .env.example app.env`.
2. Run `make build` to build the docker containers.
3. Run `make up` to run docker containers.
4. Dockerfile will automatically migrate database.
5. visit `http://locahost:8080`.

## Route List

For available router list visit
[POSTMAN Collections](https://documenter.getpostman.com/view/7576090/VUqpscbH)