FROM alpine:3.7
MAINTAINER Luca Paterlini "paterlini.luca@gmail.com"
RUN apk update && apk add python3
RUN apk add --no-cache bash
ADD ./bin /opt
WORKDIR /opt
CMD ["./main"]