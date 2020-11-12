# Argruments from FROM
ARG REGISTRY
ARG WATCHER_VERSION=2.0

# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-nodejs

ARG handler_file=handler.js
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/handler
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}

COPY ./src/handler.js .

RUN cp -r /openfx/executor/* ${HANDLER_DIR}

EXPOSE 50051

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["node"]
CMD ["fx_server.js"]
