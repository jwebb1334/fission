FROM alpine:3.18
RUN apk add --update ca-certificates
COPY . /
RUN cd /cmd/fission-bundle && chmod 0777 fission-bundle && ls -l && file fission-bundle
# RUN ls -l && file /cmd/fission-bundle/fission-bundle
ENTRYPOINT ["/cmd/fission-bundle/fission-bundle"]

