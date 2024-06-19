<div align="center">
<h1>codeanalyze</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

codeanalyze provides an opinionated perspective on top of popular static analysis capabilities such as [Semgrep](https://semgrep.dev/) to provide visibility into vulnerabilities and misconfigurations that may exist in a team's code base. Designed with data-modeling and data-integration needs in mind, codeanalyze can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The types of scans that codeanalyze can conduct are constantly growing. For the most up to date listing, please see the documentation [here](./docs/index.md)

To learn more about codeanalyze, please see the [Documentation site](https://method-security.github.io/codeanalyze/) for the most detailed information.

## Quick Start

### Get codeanalyze

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/codeanalyze`
- `docker run ghcr.io/method-security/codeanalyze`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/codeanalyze/releases/latest) page
- [Installation documentation](./getting-started/installation.md)

### General Usage

```bash
codeanalyze semgrep --config-type template --config-value <value> --target /path/to/target --local-rules-dir /path/to/rules
```

## Contributing

Interested in contributing to codeanalyze? Please see our organization wide [Contribution](https://method-security.github.io/community/contribute/discussions.html) page.

## Want More?

If you're looking for an easy way to tie codeanalyze into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, visit us [here](https://method.security)

## Community

codeanalyze is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](https://github.com/Method-Security) or our organization wide documentation [here](https://method-security.github.io).

Have an idea for a Tool to contribute? Open a Discussion [here](https://github.com/Method-Security/Method-Security.github.io/discussions).

[verify]: https://github.com/Method-Security/codeanalyze/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/codeanalyze/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/codeanalyze
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/codeanalyze
[release]: https://github.com/Method-Security/codeanalyze/releases
[releases]: https://github.com/Method-Security/codeanalyze/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/codeanalyze.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/codeanalyze/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/codeanalyze?logo=docker&label=docker%20pulls%20%2F%20codeanalyze
[docker-pull]: https://hub.docker.com/r/methodsecurity/codeanalyze
[license]: https://github.com/Method-Security/codeanalyze/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
