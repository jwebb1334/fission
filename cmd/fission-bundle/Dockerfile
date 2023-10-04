FROM alpine:3.18
RUN apk add --update ca-certificates
COPY fission-bundle /
RUN chmod +x /fission-bundle
RUN ls -l / && file /fission-bundle
ENTRYPOINT ["/fission-bundle"]

