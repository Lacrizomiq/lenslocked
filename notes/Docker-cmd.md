# Docker commands

## Make docker work

Create and start containers

```shell
docker compose up
```

## Remove containers

Stop and remove containers, networks

```shell
docker compose down
```

## Open psql CLI in Docker

```shell
docker exec -it lenslocked-db-1 /usr/bin/psql -U baloo -d lenslocked
```
