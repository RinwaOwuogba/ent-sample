### Updating Schema

- update schema
- generate updated code for db operation, in `project_root/db/ent`, run:

```
go generate
```

### Migration

- Run

```
go run migrate.go [migration name]
```

in the `project_root/db/`

- Apply migration using `golang-migrate`. in `project_root/db/` run:

```
migrate -source 'file://migrations' -database 'postgresql://postgres:postgres@127.0.0.1/test_db?sslmode=disable' up
```

### References

#### Getting started with Ent:

- https://entgo.io/docs/getting-started

#### Migration:

- https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
- https://entgo.io/blog/2022/03/14/announcing-versioned-migrations/
