# Argruments from FROM
ARG REGISTRY=keti.asuscomm.com:5000
ARG WATCHER_VERSION=0.1.0

# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-cpp

ARG handler_file=handler.cc
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/executor
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

WORKDIR ${HANDLER_DIR}
COPY ./src/handler.cc .
COPY ./src/handler.h .
RUN g++ -std=c++11 `pkg-config --cflags protobuf grpc` -c -Wall -Werror -fpic handler.cc

##################################
# Compile handler by static libray
##################################
RUN ar rcs libhandler.a handler.o

##################################
# Compile handler by shared libray
##################################
#RUN g++ -std=c++11 `pkg-config --cflags protobuf grpc` -shared -o libhandler.so handler.o

######################################
# Copy static handler libray into path
######################################
RUN cp libhandler.a /usr/local/lib

######################################
# Copy static handler libray into path
######################################
#RUN cp libhandler.so /usr/local/lib

RUN g++ -std=c++11 `pkg-config --cflags protobuf grpc`  -c -o server.o server.cc -lhandler
RUN make

ENV LD_LIBRARY_PATH=/usr/local/lib:/usr/lib:$LD_LIBRARY_PATH

EXPOSE 50051
EXPOSE 50052

HEALTHCHECK --interval=2s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./fxwatcher"]
