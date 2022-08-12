FROM golang

WORKDIR /usr/app

COPY ./ ./

RUN go get
RUN go get github.com/go-swagger/go-swagger

CMD ["swagger", "generate", "spec", "-o", "./swagger.yaml", "--scan-models"]