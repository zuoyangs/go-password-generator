FROM registry.cn-hangzhou.aliyuncs.com/zuoyang/golang:1.23.4-alpine3.2 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

RUN sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories \
  && apk add --no-cache git \
  && adduser -D golang

WORKDIR /home/golang

COPY . /home/golang/

RUN go mod init go-password-generator-api \
  && go mod tidy \
  && go build -o go-password-generator-api main.go


FROM registry.cn-hangzhou.aliyuncs.com/zuoyang/golang:1.23.4-alpine3.2

RUN adduser -D golang

WORKDIR /home/golang

COPY --from=builder /home/golang/go-password-generator-api .

RUN chown -R golang:golang /home/golang

EXPOSE 3005

USER golang

ENTRYPOINT ["./go-password-generator-api"]