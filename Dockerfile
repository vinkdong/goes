FROM alpine

RUN \
   apk update && \
   apk add curl wget ca-certificates

ADD build/bin/server /server

RUN \
  chmod a+x /server

CMD ["/server"]