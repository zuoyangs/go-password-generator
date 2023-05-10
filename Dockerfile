FROM golang:1.20.4-alpine3.17

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release 

RUN apk add --no-cache git
RUN adduser -D appuser

WORKDIR /home/appuse

COPY . .

RUN go mod init go-password-generator \
  && go mod tidy \
  && go build -o go-password-generator \
  && chown appuser:appuser go-password-generator

EXPOSE 3005

ENTRYPOINT ["./go-password-generator"]
