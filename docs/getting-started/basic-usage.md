# Basic Usage

## Binaries

Running codeanalyze as a binary allows you to more easily scan local files on your machine without worrying about Docker volumes.

```bash
codeanalyze semgrep --config-type template --config-value <value> --target /path/to/target --local-rules-dir /path/to/rules
```

## Docker

Running codeanalyze within a Docker container should typically work similarly to running directly on a host, however, but there are a few things to keep in mind.

If you're running on a Docker container on a MacOS machine and you are trying to scan a locally running service, you can leverage the `host.docker.internal` address as mentioned in the Docker documentation [here](https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host).

```bash
docker \
  -v /path/to/target:/opt/target \
  -v /path/to/rules:/opt/rules \
  methodsecurity/codeanalyze \
  semgrep \
  --config-type template \
  --config-value <value> \
  --target /opt/target \
  --local-rules-dir /opt/rules
```
