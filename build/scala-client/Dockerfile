FROM ubuntu:24.04

ARG SCALA_VERSION=2.13.15
ARG SBT_VERSION=1.10.7

LABEL org.opencontainers.image.authors="G-Research Open-Source Software"
LABEL org.opencontainers.image.licenses="Apache-2.0"
LABEL org.opencontainers.image.ref.name="Ubuntu Scala Image"
LABEL org.opencontainers.image.version=""

RUN set -ex && \
    apt update && \
    apt install -y apt-utils && \
    apt install -y bash curl && \
    curl -s -O https://downloads.lightbend.com/scala/${SCALA_VERSION}/scala-${SCALA_VERSION}.deb && \
    apt install -y ./scala-${SCALA_VERSION}.deb && \
    curl -s -L -O https://dlcdn.apache.org/maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.tar.gz && \
    tar -C / -xzvf apache-maven-*-bin.tar.gz && \
    mv /apache-maven-*/ /mvn && \
    apt-get clean && \
    rm -rf scala-${SCALA_VERSION}.deb apache-maven-*-bin.tar.gz /var/lib/apt/lists/*

ENV PATH="/mvn/bin:$PATH"
COPY scripts/build-scala-client.sh /
RUN chmod +x /build-scala-client.sh

ENTRYPOINT [ "/build-scala-client.sh" ]
