FROM golang:1.24.5-alpine
LABEL maintainer="jannes.fleming@nearxgroup.com"
ENV GO_PATH=/go/src/app
RUN apk add --no-cache git curl
WORKDIR $GO_PATH
COPY . .
RUN go mod download
RUN go install github.com/spf13/cobra-cli@latest
RUN go build -o todo-api .
RUN chmod +x ./todo-api

ENTRYPOINT ["/go/src/app/todo-api"]