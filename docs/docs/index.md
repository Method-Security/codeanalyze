# Capabilities

codeanalyze offers a variety of static code analysis tools and techniques that allow security teams to bring static code analysis into their automated workflows. Each of the below pages offers you an in depth look at a codeanalyze capability related to a static code analysis technique.

- [Semgrep](./semgrep.md)

## Top Level Flags

codeanalyze has several top level flags that can be used on any subcommand. These include:

```bash
Flags:
  -h, --help                 help for codeanalyze
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```

## Version Command

Run `codeanalyze version` to get the exact version information for your binary

## Output Formats

For more information on the various output formats that are supported by codeanalyze, see the [Output Formats](https://method-security.github.io/docs/output.html) page in our organization wide documentation.
