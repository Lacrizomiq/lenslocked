# Docker commands

## Make docker work

```shell
docker compose up
```

## Close Docker

```shell
docker compose down
```

## Open psql CLI in Docker

```shell
docker exec -it lenslocked-db-1 /usr/bin/psql -U baloo -d lenslocked
```
