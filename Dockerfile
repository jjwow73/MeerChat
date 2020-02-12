FROM golang:latest
USER root
WORKDIR /go/src/github.com/jjwow73/MeerChat
COPY . .
RUN ["go", "get", "github.com/gorilla/websocket"]
RUN ["chmod", "+x", "server.sh"]
CMD ["./server.sh"]