# Deployment

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Verification](#verification)
- [Maintenance](#maintenance)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

This guide will walk you through the steps to deploy the Golang microservice application using Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone the repository and navigate to the root directory of the project:

```bash
git clone https://github.com/samannsr/vending-machine-control-system.git
cd vending-machine-control-system
```

2. Create a .env file in the `envs` directory of the project and copy the `local.env` environment variables.

3. Run the following command to build and start the containers:

```bash
docker-compose up -d --build
```

This will build and start the following containers:

- `app`: Vending machine application

## Verification

To verify that the containers are running, use the following command:

```bash
docker-compose ps
```

## Maintenance

To stop the containers, use the following command:

```bash
docker-compose stop
```

To start the containers again, use the following command:

```bash
docker-compose start
```

To remove the containers, use the following command:

```bash
docker-compose down
```
