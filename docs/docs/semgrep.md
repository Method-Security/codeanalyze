# Semgrep

The `codeanalyze semgrep` command provides an opinionated wrapping of [Semgrep](https://semgrep.dev) to facilitate tying static code analysis capabilities into your security automation workflows.

## Usage

```bash
codeanalyze semgrep --config-type template --config-value <path to template> --target <path to directory>
```

## Help Test

```bash
codeanalyze semgrep -h
Run semgrep against code directory

Usage:
  codeanalyze semgrep [flags]

Flags:
      --config-type string       SAST config type (direct|template), direct to write custom config string (e.g. --config p/secrets), template to use a pre defined built-in and custom rule set combo
      --config-value string      SAST config value, either a string to be passed directly to semgrep CLI or a template value (e.g. secrets)
  -h, --help                     help for semgrep
      --local-rules-dir string   Absolute path to local semgrep rules directory (default "/opt/method/codeanalyze/var/conf/resources/semgrep/")
      --target string            Local folder or file code target to scan

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```
