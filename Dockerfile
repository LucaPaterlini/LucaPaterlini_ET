FROM alpine:3.7
MAINTAINER Luca Paterlini "paterlini.luca@gmail.com"
RUN apk update && apk add python3
ADD bin/main /
CMD ["/main"]