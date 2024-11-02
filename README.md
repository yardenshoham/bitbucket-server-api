# Bitbucket Server API

Go client library for the Bitbucket Server (Bitbucket Data Center) API. Generated from the OpenAPI specification.

The Atlassian Bitbucket REST API documentation can be found [here](https://developer.atlassian.com/server/bitbucket/rest/v903/intro/).

## Usage

### Install

```bash
go get github.com/yardenshoham/bitbucket-server-api
```

### Import

```go
import "github.com/yardenshoham/bitbucket-server-api/bitbucket"
```

## Development

### Generate

```bash
go generate ./...
```

### Test

We test by spinning up a local Bitbucket Data Center instance and running the tests against it.

#### Start Bitbucket Data Center

```bash
docker-compose up -d
```

This will start a local Bitbucket Data Center instance on port 7990. We must manually register a user with the following credentials:

- username: `admin`
- password: `admin`

You need an Atlassian account to get a trial license key.

#### Run tests

Once the Bitbucket Data Center instance is running and the user is registered, you can run the tests:

```bash
go test ./...
```