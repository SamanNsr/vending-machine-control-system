# Design

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Key Components and Features](#key-components-and-features)
- [Design Decisions](#design-decisions)
  - [See also](#see-also)
- [Protocol Buffer](#protocol-buffer)
- [API docs](#api-docs)
- [Layout](#layout)
- [Error Handling](#error-handling)
- [Diagrams and Mockups](#diagrams-and-mockups)
- [Open Issues and Areas for Improvement](#open-issues-and-areas-for-improvement)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

This document outlines the design of the project, including the key components and features, design decisions, and technologies used.

## Key Components and Features

The project includes the following components and features:

- A PostgreSQL database for storing and querying data
- A Kafka message broker for reliable, scalable event streaming
- A Redis cache for improving performance and reducing load on the database
- A gRPC API for efficient communication between services
- An Echo framework for building HTTP APIs and web applications
- Sentry integration for error tracking and reporting

## Design Decisions

The following design decisions were made:

- PostgreSQL was chosen as the primary data store because of its robust support for ACID transactions and ability to handle large volumes of data.
- Kafka was chosen as the message broker because of its high performance and ability to handle large amounts of data.
- Redis was chosen as the cache because of its in-memory storage and ability to support multiple data structures.
- gRPC was chosen as the API technology because of its efficient binary encoding and ability to support streaming requests and responses.
- Echo was chosen as the web framework because of its lightweight and easy-to-use design.
- Sentry was chosen for error tracking and reporting because of its comprehensive feature set and integrations with a wide range of technologies.
- The project follows a clean architecture design pattern, with separate layers for the domain logic, application logic, and infrastructure. - This helps to improve the maintainability and testability of the codebase.

### See also

- [Echo](https://echo.labstack.com/) for web server routing
- [Zap Logger](https://github.com/uber-go/zap) for logging
- [OZZO](github.com/go-ozzo/ozzo-validation) for data validation


## API docs

The template doesn't have API docs. For auto-generated API docs that you include, you can also give instructions on the
build process.

## Layout

The project is organized into a number of directories and subdirectories, as shown in the following tree structure:

```tree
├── .github
│  └── pull_request_template.md
├── app
│  └── app.go
├── cmd
│  └── main.go
├── envs
│  ├── .env
│  ├── local.env
│  ├── production.env
│  ├── stage.env
│  └── test.env
├── internal
│  └── module_name
│    ├── configurator
│    ├── delivery
│    │  └── http
│    ├── domain
│    ├── dto
│    ├── exception
│    ├── repository
│    ├── tests
│    │  ├── fixtures
│    │  └── integrations
│    └── usecase
├── pkg
│  ├── config
│  ├── constant
│  ├── env
│  ├── error
│  │  ├── custom_error
│  │  ├── error_utils
│  │  └── http
│  ├── http
│  ├── infrastructure
│  └── logger
│
├── .gitignore
├── .pre-commit-config.yaml
├── golangci.yaml
├── docker-compose.e2e-local.yaml
├── docker-compose.yaml
├── go.sum
├── Makefile
├── go.mod
├── LICENSE
└── README.md
```

- `.github`: Contains GitHub-specific files, such as CODEOWNERS and pull request template.
- `app`: Contains the entry point of the application.
- `cmd:` Contains the main command of the application.
- `envs`: Contains the environment configuration files.
- `internal`: Contains the internal modules.
- `pkg`: Contains the shared packages.
- `.gitignore`: Defines which files and directories should be ignored by Git.
- `.pre-commit-config.yaml`: Defines the pre-commit hooks and checks.
- `golangci.yaml`: Defines the configuration for the GolangCI linter.
- `docker-compose.yaml`: Defines the configuration for the Docker Compose containers.
- `go.sum`: Contains the checksums for the Go module dependencies.
- `Makefile`: Contains the build and development tasks.
- `go.mod`: Defines the Go module dependencies.
- `LICENSE`: Contains the license terms for the project.
- `README.md`: Contains the project documentation.

## Error Handling

The project includes a custom error implementation to handle errors in a consistent and structured way. The custom error implementation includes a set of predefined error types for different categories of errors, such as  bad request errors, internal server errors, and so on. This allows us to classify and handle errors in a standardized way, making it easier to understand and debug issues that may arise.

The custom error implementation is located in the error package in the pkg directory. It includes the following error types:

- `BadRequestError`: Represents an error that occurs when the request is invalid or malformed, such as a missing or invalid parameter.
- `InternalError`: Represents an error that occurs within the infrastructure layer, such as a failure to connect to a database or a failure to access a required service.
- `NotFoundError`: Represents an error that occurs when the requested resource is not found.
- `ValidationError`: Represents an error that occurs when the request data is invalid, such as when a required field is missing or a field is in the wrong format.

The custom error implementation also includes utility functions for creating and wrapping errors, as well as functions for handling and responding to errors in different contexts, such as when handling HTTP requests or when working with gRPC.

To use the custom error implementation, the application code can import the error package and use the error types and utility functions as needed. For example, to create a new BadRequestError, the code can use the NewBadRequestError function:

```golang
import "project/pkg/error/custum_errors"

err := customErrors.NewBadRequestError("invalid parameter", code)
```

To wrap an existing error with a custom error type, the code can use the WrapError function:

```golang
import "project/pkg/error/custum_errors"

err := customErrors.NewBadRequestErrorWrap(err, "invalid parameter", code)
```

## Diagrams and Mockups

To include diagrams or mockups in the design.md file to illustrate the design of your project, you will need to create these visualizations using a diagramming tool or software. Here are a few options for creating diagrams and mockups:

- dbdiagram.io: an online tool for creating and sharing database diagrams. It allows you to design the schema of your database visually, using a simple drag-and-drop interface. You can add tables, columns, and relationships to the diagram, and customize the appearance of the elements using a range of formatting options.
- Draw.io: A web-based diagramming tool that allows you to create a wide range of diagrams, including flowcharts, mind maps, and UML diagrams.
- Lucidchart: A web-based diagramming tool that offers a range of templates and shapes for creating professional-quality diagrams.
- Figma: A web-based design and prototyping tool that allows you to create wireframes, mockups, and prototypes for web and mobile applications.

Once you have created the diagrams or mockups that you want to include in the design.md file, you can add them to the file by including the images inline or by linking to the images.

[Include any relevant diagrams or mockups to illustrate the design]

## Open Issues and Areas for Improvement

[Describe any open issues or areas for improvement in the design]
