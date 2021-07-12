# Argruments from FROM
ARG REGISTRY=10.0.0.170:5000
ARG WATCHER_VERSION=2.0

# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-cpp

ARG handler_file=handler.cc
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/executor
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

RUN mkdir -p ${HANDLER_DIR}/src
WORKDIR ${HANDLER_DIR}/src
COPY ./src/handler.cc .
COPY ./src/handler.h .
#COPY ./src/CMakeLists.txt .

WORKDIR ${HANDLER_DIR}
#COPY ./src/handler.cc .
#COPY ./src/handler.h .
COPY cmake cmake/
COPY CMakeLists.txt .
RUN mkdir -p cmake/build
WORKDIR cmake/build
RUN cmake -DCMAKE_PREFIX_PATH=$MY_INSTALL_DIR ../..
RUN make -j

EXPOSE 50051
EXPOSE 50052

HEALTHCHECK --interval=5s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./server"]
