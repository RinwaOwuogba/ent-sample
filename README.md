### Updating Schema

- update schema
- generate updated code for db operations, run:

```
go generate ./db/ent
```

### Migration

- Create migration files by running:

```
go run db/make-migrations.go [migration name]
```

- Apply migration using `golang-migrate`. in project_root run:

```
migrate -source 'file://db/migrations' -database 'postgresql://postgres:postgres@127.0.0.1/test_db?sslmode=disable' up
```

### References

#### Getting started with Ent:

- https://entgo.io/docs/getting-started

#### Migration:

- https://entgo.io/blog/2022/03/14/announcing-versioned-migrations/
- https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
