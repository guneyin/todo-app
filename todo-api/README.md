
![Coverage](https://img.shields.io/badge/Coverage-66.7%25-yellow)


# todo-api service

## API Reference

#### List all todos

```http
  GET /api/todos
```

Sample response
```json
{
    "todos": [
      "buy some milk",
      "pay bills"
    ]
}
```

#### Create a todo

```http
  POST /api/todos
```

Sample request
```json
{
    "todo": "buy some milk"
}
```

## Environment Variables

| Key      | Value  | Mandatory | Sample |
|----------|--------|-----------|:-------|
| API_PORT | string | yes       | 3001   |

## Test

### Unit tests
```bash
  go test ./...
```

### E2E tests
```bash
  go test -v . -run Test_E2E
```

## Project Architecture

- Project based on almost Golang standard library. No any external http framework used. Project build on service layout, there are two layers, handler and service.
  - Handler layer: Responsible to handle REST requests
  - Service layer: Stores todo list state, process requests.

- Database: In-memory store used to keep project simple
- Linter: golangci used for it's simplicity and wide configuration options
- Build: Makefile for it's simplicity
- Testing framework: Testify, lightweight and easy to implement
- E2E Testing: Contract based testing with pact files


    