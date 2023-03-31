# Devstore

## docker compose

1. Running : `docker compose up`
2. Running as daemon : `docker compose up -d`
3. Check running container : `docker ps`
4. Stop container `docker compose stop`
5. Stop & remove conainer `docker compose down`

## Contianer Interaction

### Database

1. enter to db container : `docker exec -it devstore-db-1 bash`
2. enter to postgres console : `psql -U postgres -d postgres`

### Migration

1. Create : `docker compose -f docker-compose.yaml --profile tools run --rm migrate create -ext sql -dir /migrations create_table_category`
2. Up : `docker compose -f docker-compose.yaml --profile tools run --rm migrate up`
3. Down : `docker compose -f docker-compose.yaml --profile tools run --rm migrate down`
