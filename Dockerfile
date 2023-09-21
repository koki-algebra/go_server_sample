FROM golang:1.21.1 AS builder

ARG PROTOCOL=grpc

ENV ROOT=/go/src

COPY . ${ROOT}

WORKDIR ${ROOT}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/${PROTOCOL}/main.go


FROM alpine:3.18.3 AS deploy

RUN apk add --no-cache tzdata

ENV ROOT=/go/src

COPY --from=builder ${ROOT}/main ${ROOT}/

CMD [ "/go/src/main" ]
