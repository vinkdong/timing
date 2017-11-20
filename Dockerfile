FROM alpine:latest
RUN \
   mkdir timingconf
COPY timing /
COPY config.yaml /timingconf/
EXPOSE 9980
CMD [ "./timing","--conf","/timingconf/config.yaml","--enable_metrics" ]