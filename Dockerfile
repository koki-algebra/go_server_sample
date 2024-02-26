FROM golang:1.22.0 AS builder

ARG PROTOCOL=grpc

ENV ROOT=/go/src

WORKDIR ${ROOT}

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/${PROTOCOL}


FROM alpine:3.18.3 AS deploy

ENV ROOT=/go/src

WORKDIR ${ROOT}

RUN apk add --no-cache tzdata

COPY --from=builder ${ROOT}/main .

CMD [ "/go/src/main" ]
