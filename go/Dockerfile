# Argrumnets for FROM
ARG REGISTRY
ARG GO_VERSION=1.9.7
ARG WATCHER_VERSION=0.1.0

# Get fxwatcher - if fxwatcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-go as watcher

ARG handler_file
ARG handler_name=Handler

ENV HANDLER_DIR=${GOPATH}/src/openfx/handler
ENV HANDLER_FILE=${GOPATH}/src/openfx/handler/handler
ENV HANDLER_NAME=${handler_name}

RUN cp ${GOPATH}/src/github.com/keti-openfx/openfx-watcher/go/fxwatcher /fxwatcher

RUN chmod +x /fxwatcher

RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}
COPY ./src/handler.go .
COPY ./Gopkg.toml .
RUN dep ensure
RUN go build -o ${HANDLER_FILE} -buildmode=plugin .

HEALTHCHECK --interval=2s CMD [ -e /tmp/.lock ] || exit 1

CMD ["/fxwatcher"]
