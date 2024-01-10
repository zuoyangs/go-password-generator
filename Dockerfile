FROM registry.cn-hangzhou.aliyuncs.com/mw5uk4snmsc/go:1.21.5-alpine3.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

RUN sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories \
  && apk add --no-cache git \
  && adduser -D appuser

WORKDIR /home/appuser

COPY . /home/appuser/

RUN go mod init go-password-generator \
  && go mod tidy \
  && go build -o go-password-generator main.go


FROM registry.cn-hangzhou.aliyuncs.com/mw5uk4snmsc/go:1.21.5-alpine3.19

RUN adduser -D appuser \
  && mkdir -pv /home/appuser/static

WORKDIR /home/appuser

COPY static /home/appuser/static/
COPY --from=builder /home/appuser/go-password-generator .

RUN chown -R appuser:appuser /home/appuser

EXPOSE 3005

USER appuser

ENTRYPOINT ["./go-password-generator"]
