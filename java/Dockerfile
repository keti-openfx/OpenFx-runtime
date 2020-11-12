# Argruments from FROM
ARG REGISTRY=keti.asuscomm.com:5000
ARG WATCHER_VERSION=2.0

FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-java

ARG handler_file=Handler.java
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/executor
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

WORKDIR ${HANDLER_DIR}
COPY ./src/Handler.java ${HANDLER_DIR}/src/main/java/io/grpc/fxwatcher

RUN ./gradlew clean build
RUN ./gradlew installDist

EXPOSE 50051

HEALTHCHECK --interval=5s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./build/install/java/bin/fxwatcher"]
