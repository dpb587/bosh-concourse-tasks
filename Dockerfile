FROM alpine:3.4
RUN true \
  && apk --no-cache add bash ca-certificates curl \
  && curl -sLo /usr/local/bin/bosh https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-2.0.1-linux-amd64 \
  && echo 'fbae71a27554be2453b103c5b149d6c182b75f5171a00d319ac9b39232e38b51  /usr/local/bin/bosh' | sha256sum -c \
  && chmod +x /usr/local/bin/bosh
ADD bin/ /usr/local/bin/
