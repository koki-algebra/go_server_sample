FROM golang:1.21.1-bullseye AS builder

ENV ROOT=/go/src
ENV BUILD_DIR=cmd/grpc

COPY . ${ROOT}

WORKDIR ${ROOT}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ${BUILD_DIR}/main.go


FROM alpine:3.18.3 AS deploy

RUN apk add --no-cache tzdata

ENV ROOT=/go/src

COPY --from=builder ${ROOT}/main ${ROOT}/

CMD [ "/go/src/main" ]
