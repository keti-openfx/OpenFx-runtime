# Argruments from FROM
ARG REGISTRY
ARG WATCHER_VERSION=2.0

# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-python3

ARG handler_file=handler.py
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/handler
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}
ENV PYTHONUNBUFFERED=0

RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}
COPY ./src/handler.py .
COPY ./requirements.txt .
RUN touch ${HANDLER_DIR}/__init__.py
RUN pip3 install -r requirements.txt

RUN cp -r /openfx/executor/* ${HANDLER_DIR}

EXPOSE 50051
EXPOSE 50052

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["python3"]
CMD ["server.py"]
