# `viter` - Striboh's core service

## Development

### Setup local Postgres database

Repo has `docker-compose.yml` configuration for setting up local
Postgres DB by:

```shell
docker compose up -d
```

### Starting up

> [!WARNING]
>
> Before starting up, make sure that all configured parameters (see
> all available environmental variables in
> [`config.go`](./internal/config/config.go)) are not specified.

To run the service:

```shell
go run main.go
```

> [!NOTE] 
>
> `viter` has optional flag `-migrate` which specifies what
> strategy should service use: apply migrations `-migrate=up`, revert
> them `-migrate=down` or do nothing (don't specify `-migrate`).

### Code generation from OpenAPI spec

For generation of code from `openapi` spec
([see](./openapi/openapi.yaml)), install
[`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen), and
then after updating `openapi.yaml`:

```shell
go generate
```

should work fine and update the `internal/service/api/api.gen.go`.
