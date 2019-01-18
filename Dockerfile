FROM golang:1.11.4-alpine3.8 as builder
WORKDIR /go/src/github.com/vinkdong/timing
ADD . .

RUN \
  apk add gcc build-base
RUN \
  go build .


FROM alpine:latest
RUN \
   mkdir timingconf && \
   apk add libc6-compat

RUN \
set -ex \
   && apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/vinkdong/timing/timing /timing

COPY config.yaml /etc/config.yml
EXPOSE 9980
CMD [ "./timing","--conf","/etc/config.yml","--enable_metrics" ]