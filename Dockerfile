FROM golang:alpine as step-s2
RUN apk update && apk add --no-cache bash git openssh
RUN go get -u github.com/golang/dep/...
WORKDIR /go/src/github.com/cocoagaurav/devops
COPY . ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go install /go/src/github.com/cocoagaurav/devops/cmd/s2


FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=step-s2 /go/bin/s2 /bin/s2
EXPOSE 8007
ENTRYPOINT ["s2"]
CMD ["-logtostderr"]