FROM golang

RUN  apt update && \
     apt install -y apt-transport-https curl gnupg && \
     curl "https://repo.zelenin.pw/gpg.key" |  apt-key add - &&\
     echo "deb [arch=amd64] https://repo.zelenin.pw common contrib" |  tee "/etc/apt/sources.list.d/tdlib.list" && \
     apt update &&\
     apt install -y tdlib-dev

COPY ./ /go/src/github.com/zoh/telega-fwd-golang
WORKDIR /go/src/github.com/zoh/telega-fwd-golang

RUN go get ./...
RUN go build -o telega-fwd cmd/main.go

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD ./telega-fwd
