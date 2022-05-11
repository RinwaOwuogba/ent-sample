### Steps

- Update schema, generate schema code

### Migration

- Run

```
go run migrate.go [migration name]
```

in the db directory

- Apply migration using `golang-migrate`.

```
migrate -source 'file://migrations' -database 'postgresql://postgres:postgres@127.0.0.1/test_db?sslmode=disable' up
```

### More References

#### Getting started:

- https://entgo.io/docs/versioned-migrations/#from-graph

#### Migration:

- https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
- https://entgo.io/blog/2022/03/14/announcing-versioned-migrations/
