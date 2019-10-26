FROM scratch

ARG buildSha
ENV BUILD_SHA=$buildSha

WORKDIR /app

COPY ms-template service
COPY metadata.json metadata.json

ENTRYPOINT ["/app/service"]