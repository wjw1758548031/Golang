FROM alpine:3.6
LABEL author="xqychj(xqychj@163.com)"


WORKDIR /usr/src/app

COPY ["./" , "./"]

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN chmod 777 -R /usr/src/app
EXPOSE 10002

CMD ["./haikangplate"]
