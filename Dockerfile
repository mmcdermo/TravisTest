FROM golang

RUN go get "github.com/gorilla/mux" && \
    go get "github.com/gorilla/handlers" && \
    go get "github.com/lib/pq" && \
    go get "gopkg.in/yaml.v2"


ENV GOBIN /go/bin/
ENV GOPATH /go/
ENV GOSRC /go/src/
ENV PATH $GOBIN:$PATH:$GOSRC

RUN mkdir /go/src/TravisTest/
WORKDIR /go/src/TravisTest/

COPY ./ /go/src/TravisTest/

#COPY ./docker-entrypoint.sh /docker-entrypoint.sh
#RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/go/bin/main"]


WORKDIR /go/src/TravisTest
RUN go install /go/src/TravisTest/main.go

EXPOSE 80
EXPOSE 443
