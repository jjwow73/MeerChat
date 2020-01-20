FROM golang:latest
USER root
WORKDIR /go/src/github.com/wkd3475/MeerChat
COPY . .
RUN ["chmod", "+x", "meerchat.sh"]
CMD ["./meerchat.sh"]