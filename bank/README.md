# Bank Application

## Postgres
```bash
docker logs postgres15
docker exec -it postgres15 bash
docker exec -it postgres15 psql -U root
docker exec -it postgres15 psql -U root simple_bank
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

## Dependencies
```bash
brew install golang-migrate sqlc
```

### sqlc
Generate `sqlc.yaml`
```bash
sqlc init
mkdir db/sqlc
mkdir db/query
```
