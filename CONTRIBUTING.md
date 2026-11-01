# Contributing

Thank you for helping develop Mayanna–HashMin. This is an experimental research project. Contributions may include paper corrections, cryptographic analysis, code, tests, documentation and benchmarks.

## Before starting

Read the [README](README.md), [paper](docs/Mayanna-HashMin-Primitives.pdf) and [security policy](SECURITY.md). Search existing issues and pull requests. For substantial changes, open an issue describing the problem, proposed approach and relationship to the paper before investing in implementation.

## Development

Use Go 1.22.2 or later. Fork the repository, create a branch and keep changes focused.

```sh
go test ./...
go vet ./...
```

Format changed Go files with `gofmt`. Add meaningful tests for changed behaviour. Cryptographic changes should include rejected/forged proof cases and an explanation of assumptions, transcript binding, randomness and serialization where relevant. Passing tests alone does not establish cryptographic security.

## Pull requests

Explain the problem, the resulting behaviour, validation performed and remaining limitations. Link relevant issues and paper sections. Separate implemented behaviour from proposed guarantees. Avoid adding real credentials, private keys, personal data or confidential reports to examples or tests.

## Licensing and conduct

By submitting a contribution, you agree that your contribution is provided under the repository's [MIT License](LICENSE). Preserve third-party attribution and only submit material you have the right to contribute. Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
