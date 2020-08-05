FROM golang:latest

WORKDIR /seekjob
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build" --command="./seekjob"