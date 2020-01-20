FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD hello-web /
ENTRYPOINT ["/hello-web"]
