# Migrations

## Usage

1. Install package [github.com/rubenv/sql-migrate](https://github.com/rubenv/sql-migrate) globaly

```shell
go get -v github.com/rubenv/sql-migrate/...
```

2. Go to the migrations folder in your package and generate migrations file like described [here](https://github.com/rubenv/sql-migrate/blob/master/README.md).

```shell
sql-migrate new create_users_table
```


3. Write SQL in generated migration files.
