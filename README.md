# Go Postgres Connection Test application

This is a simple application to test the connection to a Postgres database.

## Deploy

```bash
helm upgrade --install db-connection-test ./infra/go-db-connection-test --set db.connectionString="postgres://postgres:secretpw1@postgres-postgresql:5432/postgres"
```

## Usage

To test the connection, run the following command:

```bash
curl http://localhost:8080/connect
```
