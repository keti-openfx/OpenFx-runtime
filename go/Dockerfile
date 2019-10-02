# Argrumnets for FROM
ARG REGISTRY
ARG GO_VERSION=1.9.7
ARG WATCHER_VERSION=0.1.0

# Get fxwatcher - if fxwatcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-go as watcher

ARG handler_file
ARG handler_name=Handler

ENV HANDLER_DIR=${GOPATH}/src/github.com/keti-openfx/openfx/executor/go/handler
ENV HANDLER_FILE=${HANDLER_DIR}/handler
ENV HANDLER_NAME=${handler_name}

RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}
COPY ./src/handler.go .

WORKDIR ${GOPATH}/src/github.com/keti-openfx/openfx/executor/go
RUN go build -o ${HANDLER_FILE} -buildmode=plugin ${HANDLER_DIR}/handler.go

EXPOSE 50051

HEALTHCHECK --interval=2s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./fxwatcher"]
