# vending-machine-control-system

# Go Microservice Template

This is the code repository for [Vending Machine System Control](https://www.github.com/samannsr/vending-machine-control-system), simple API to handle vending machines.

**Build, test, and deploy robust golang microservices**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About the project](#about-the-project)
- [Status](#status)
- [Reference](#reference)
- [Getting started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Setup](#setup)
- [Testing](#testing)
    - [Linting](#linting)
    - [Building](#building)
    - [Cleaning](#cleaning)
    - [Continuous Integration](#continuous-integration)
- [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## About the project

The repo is used to handle vending machine as a central systems.

## Reference

- [Deployment Instructions](docs/deployment.md)
- [Design Decisions and Technical Considerations](docs/design.md)
- [Environment Variables](docs/env.md)

## Getting started

### Prerequisites

- [Go 1.13+](https://golang.org/doc/install)

### Setup

1. Clone the repository and navigate to the root directory of the project:

```bash
git clone https://github.com/samannsr/vending-machine-control-system.git
cd vending-machine-control-system
```

2. Install the dependencies:

```bash
make dep
```

3. Run the development server:

```bash
make run_dev
```

This will start the http server on port 8080. You can customize the ports by setting the `HTTP_PORT` environment variable.

## Testing

To run the tests, use the following command:

```bash
make test
```

To generate a test coverage report, use the following command:

```bash
make test_coverage
```

### Linting

To lint the code, use the following command:

```bash
make lint
```

This will run all available linters, including Go lint, Dockerfile lint, and YAML lint.

### Building

To build the binary, use the following command:

```bash
make build
```

This will create a binary in the `out/bin` directory. You can run the binary with the following command:

```bash
make run
```

### Cleaning

To clean the build artifacts and generated files, use the following command:

```bash
make clean
```

This will remove the `bin` and `out` directories, as well as any build-related files.

## Notes
