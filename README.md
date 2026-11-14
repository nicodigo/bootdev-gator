# Gator

## How to run

1. Create a ~/.gatorconfig.json with:
```bash
{"db_url":"postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@localhost:5432/<POSTGRES_DB>","current_user_name":""}
```
2. Copy example.env to .env and edit the enviroment variables.
```bash
cp example.env .env
```
3. Start the postgres database
```bash
docker compose up
```
4. Create the migrations
```bash
cd sql/schema
goose postgres 'postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@localhost:5432/<POSTGRES_DB>' up
```
5. Build the binary
```bash
go build -o gator
```
6. Register an user
```bash
./gator register <username>
```
7. Add a feed
```bash
./gator addfeed <feed_name> <feed_url>
```
8. Run the aggregator (scrapes and saves posts)
```bash
./gator agg <time_between_reqs>
```
9. Browse posts
```bash
./gator browse <limit>
```


