FROM alpine:latest
RUN \
   mkdir timingconf && \
   apk add libc6-compat

RUN \
set -ex \
   && apk add --no-cache ca-certificates

COPY timing /
COPY config.yaml /etc/config.yml
EXPOSE 9980
CMD [ "./timing","--conf","/etc/config.yml","--enable_metrics" ]