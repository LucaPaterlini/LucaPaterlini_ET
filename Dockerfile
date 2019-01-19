FROM alpine:3.7
MAINTAINER Luca Paterlini "paterlini.luca@gmail.com"
RUN apt-get update && apt-get install python3
ADD bin/main /
CMD ["/main"]