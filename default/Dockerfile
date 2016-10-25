FROM alpine

RUN apk add go=1.6.3-r0 --update-cache && \
    apk add git --update-cache && \
    mkdir /service

ENV GOPATH /service
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

COPY . /service

RUN go get github.com/vetbase/husky && \
    go build -i -o /service/bin/default /service/main.go && \
    cp /service/bin/default /usr/bin/default

EXPOSE 8080
