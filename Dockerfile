FROM golang:1.12.2-alpine3.9
RUN apk add git
ADD . /go/src/quiz-cli
WORKDIR /go/src/quiz-cli
RUN go get /go/src/quiz-cli
RUN go install
ENTRYPOINT ["/go/bin/quiz-cli"]
