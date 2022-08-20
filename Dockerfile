FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY . .

RUN apk update && apk add curl && apk add bash && apk add make

RUN go mod download

RUN echo "Downloading cosmtrek/air for live reload"

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN #migrate -path db/migrations -database "mysql://root:root@tcp(go-auth-mysql:3306)/go_db?multiStatements=true" -verbose up

ENTRYPOINT [ "sh","./scripts/entrypoint.sh" ]