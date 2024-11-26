# Agricola

Agricola is a tool for managing deployments declaratively. The goal of the tool is to be able to deploy and manage static websites and application using Docker. The project is currently heavily under development, and it is not ready for use.

## Goals

The first goal of this project is to fit my needs (that’s why I’m writing it). The goals listed here reflect that.

- Zero-downtime deployments
- Serve static websites as is (there is no need to set up a Docker container for static websites)
- Serve other applications using Docker
- Set up SSL and handle renewals

## Development

I aim to minimize the use for dependencies and use only minimal dependencies when needed. The Go standard library has most of the batteries included and there is really no need to bring in the kitchen sink.

Also, the linter rules used for now are really strict. This may be overshooting, but as the project is brand new there is no reason to limit the rules because of legacy code. The linter rules are revised as needed.

To build the project, run:

    make build

To lint the project, run:

    make check

## License

The project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more information.
