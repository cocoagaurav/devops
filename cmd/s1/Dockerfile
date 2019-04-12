#building
FROM golang:alpine as step-1
RUN apk update && apk add --no-cache bash git openssh
RUN go get -u github.com/golang/dep/...
WORKDIR /go/src/github.com/cocoagaurav/devops
COPY . ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go install /go/src/github.com/cocoagaurav/devops/cmd/s1

#main image
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=step-1 /go/bin/s1 /bin/s1
EXPOSE 8008
ENTRYPOINT ["s1"]
CMD ["-logtostderr"]