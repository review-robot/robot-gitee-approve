FROM golang:latest as BUILDER

MAINTAINER zengchen1024<chenzeng765@gmail.com>

# build binary
COPY . /go/src/github.com/opensourceways/robot-gitee-approve
RUN cd /go/src/github.com/opensourceways/robot-gitee-approve && CGO_ENABLED=1 go build -v -o ./robot-gitee-approve ./server

# copy binary config and utils
FROM golang:latest
RUN  mkdir -p /opt/app/
# overwrite config yaml
COPY  --from=BUILDER /go/src/github.com/opensourceways/robot-gitee-approve/robot-gitee-approve /opt/app

WORKDIR /opt/app/
ENTRYPOINT ["/opt/app/robot-gitee-approve"]