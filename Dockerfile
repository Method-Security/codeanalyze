FROM alpine:3.20 as base

ARG CLI_NAME="codeanalyze"

RUN \
  apk update && \
  apk add bash jq python3-dev py3-pip pipx gcc musl-dev libffi-dev && \
  pipx install semgrep

# Ensure pipx binaries are in the PATH
ENV PATH="/root/.local/bin:$PATH"

# Setup Method Directory Structure
RUN \
  mkdir -p /opt/method/${CLI_NAME}/ && \
  mkdir -p /opt/method/${CLI_NAME}/var/data && \
  mkdir -p /opt/method/${CLI_NAME}/var/data/tmp && \
  mkdir -p /opt/method/${CLI_NAME}/var/conf && \
  mkdir -p /opt/method/${CLI_NAME}/var/log && \
  mkdir -p /opt/method/${CLI_NAME}/service/bin && \
  mkdir -p /mnt/output

COPY configs/* /opt/method/${CLI_NAME}/var/conf/
COPY ${CLI_NAME} /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}

RUN \
  adduser --disabled-password --gecos '' method && \
  chown -R method:method /opt/method/${CLI_NAME}/ && \
  chown -R method:method /mnt/output

USER method

WORKDIR /opt/method/${CLI_NAME}/

ENV PATH="/opt/method/${CLI_NAME}/service/bin:${PATH}"
