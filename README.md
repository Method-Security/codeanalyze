# Codeanalyze

Codeanalyze analyzes local code repositories or directories for configurations and vulnerabilities. It can be run locally or as part of CI.

## Adding a new codeanalyze Capability

TODO update the below

1. Add a file to `cmd/` that corresponds to the sub-command name you'd like to add to the `codeanalyze` CLI
2. You can use `cmd/ec2.go` as a template
3. Your file needs to be a member function of the `AwsEnumerate` struct and should be of the form `Init<cmd>Command`
4. Add a new member to the `AwsEnumerate` struct in `cmd/root.go` that corresponsds to your command name. Remember, the first letter must be capitalized.
5. Call your `Init` function from `main.go`
6. Add logic to your commands runtime and put it in its own package within `internal` (e.g., `internal/ec2`)

## Testing

### Testing from Source (pre-build)

You can test locally without building by running

```bash
go run main.go <subcommand> <flags>
```

### Testing the CLI (post-build)

You can test locally using the CLI by building it from source. Run, `./godelw clean && ./godelw build` to clean out the `out/` directory and rebuild. You will now have a binary at `out/build/codeanalyze/<version>/darwin-arm64/codeanalyze` that you can run

## Building the Docker Container

I have not yet figured out how to get godel to build docker for us, so at the moment, it's a bit of a pain. The best idea is to follow what the `build-docker` stage in `.gitlab-ci.yml` does
