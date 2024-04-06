# go-short

basic url shortener using go, postgres and redis.

This project consists of three microservices: `id-service`, `redirection-service`, and `short-service`.

## id-service

This service is responsible for generating unique IDs. The main entry point is [`main.go`](id-service/cmd/server/main.go) in the `cmd/server` directory. The ID generation logic is implemented in [`handler.go`](id-service/pkg/api/ids/handler.go) in the `pkg/api/ids` directory.

## redirection-service

This service handles redirection from short URLs to their original URLs. The main entry point is [`main.go`](redirection-service/cmd/server/main.go) in the `cmd/server` directory. The redirection logic is implemented in [`handler.go`](redirection-service/pkg/api/links/handler.go) in the `pkg/api/links` directory.

## short-service

This service creates short URLs. The main entry point is [`main.go`](short-service/cmd/server/main.go) in the `cmd/server` directory. The short URL creation logic is implemented in [`handler.go`](short-service/pkg/api/links/handler.go) in the `pkg/api/links` directory. The database schema is defined in [`init.sql`](short-service/pkg/db/init.sql) in the `pkg/db` directory.

## Getting Started

To get started, clone the repository and install the dependencies for each service. Then, you can run each service locally.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.