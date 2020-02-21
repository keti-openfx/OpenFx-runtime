# Argruments from FROM
ARG REGISTRY
ARG WATCHER_VERSION=0.1.0

# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/fxwatcher:${WATCHER_VERSION}-csharp

ARG handler_file=handler.cs
ARG handler_name=Handler

ENV HANDLER_DIR=/openfx/handler
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}

COPY ./src/handler.cs .

RUN cp -r /openfx/executor/* ${HANDLER_DIR}

EXPOSE 50051

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

RUN dotnet publish -o out -c Release -f netcoreapp3.1 fxServer.csproj
ENTRYPOINT ["dotnet", "exec", "/openfx/handler/out/fxServer.dll"]  
