FROM alpine:3.18
RUN apk add --update ca-certificates
COPY pre-upgrade-checks /
ENTRYPOINT ["/pre-upgrade-checks"]
