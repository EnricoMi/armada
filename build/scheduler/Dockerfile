ARG BASE_IMAGE=alpine:3.21.3

FROM ${BASE_IMAGE}
LABEL org.opencontainers.image.title=scheduler
LABEL org.opencontainers.image.description="Scheduler"
LABEL org.opencontainers.image.url=https://hub.docker.com/r/gresearch/scheduler

RUN addgroup -S -g 2000 armada && adduser -S -u 1000 armada -G armada
USER armada
COPY scheduler /app/
COPY config/scheduler/config.yaml /app/config/scheduler/config.yaml
COPY config/logging.yaml /app/config/logging.yaml
WORKDIR /app
ENTRYPOINT ["./scheduler"]
