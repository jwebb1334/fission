FROM alpine:3.18
RUN apk add --update ca-certificates
COPY . .
RUN ls -l 
ENTRYPOINT ["/bin/sh", "-c", "./fission-bundle"]

