FROM scratch

ARG buildSha
ENV BUILD_SHA=$buildSha

WORKDIR /app

COPY output/service service
COPY metadata.json metadata.json

ENTRYPOINT ["/app/service"]