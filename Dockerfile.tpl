FROM alpine

RUN \
   set -ex \
   apk update &&\
   apk add bash &&\
   apk add curl wget ca-certificates

ADD {{APP_NAME}} /{{APP_NAME}}

RUN \
  chmod a+x /{{APP_NAME}}

CMD ["/{{APP_NAME}}"]