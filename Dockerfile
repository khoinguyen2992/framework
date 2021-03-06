FROM golang:1.8-alpine
MAINTAINER Khoi Nguyen <minhkhoi@siliconstraits.com>
ENV APP_HOME $GOPATH/src/framework
WORKDIR $APP_HOME
RUN apk update && apk upgrade && apk add --no-cache postgresql-client bash git openssh openjdk7-jre gcc g++
RUN go get github.com/tools/godep
RUN go get bitbucket.org/liamstask/goose/cmd/goose
COPY ./Godeps $APP_HOME/Godeps
RUN godep restore
COPY . $APP_HOME
RUN chmod 775 ./migrate_db.sh ./test.sh
EXPOSE 8080
CMD ["/bin/sh", "./docker-start.sh"]
