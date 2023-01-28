# Bank Application

## Postgres
```bash
docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres

docker exec -it some-postgres psql -U root

docker logs some-postgres
```

Run a query in the container console:
```postgresql
select now();

\q # quit
```

## Database migration
```bash
migrate create -ext sql -dir db/migration -seq init_schema
```