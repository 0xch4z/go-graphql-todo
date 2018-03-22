FROM golang:1.10-stretch

ENV SRCDIR=/go/src/github.com/charliekenney23/go-graphql-todo
WORKDIR $SRCDIR
ADD . $SRCDIR
RUN go build -o main

EXPOSE 80

ENTRYPOINT [ "./main" ]
