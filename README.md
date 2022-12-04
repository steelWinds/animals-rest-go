# animals-rest-go
Simple REST-API with use Docker, PostgreSQL and Golang (GIN & Gorm)

## Getting Started

### Configuration

Yout should create ```postgres.evn``` file with next envs:

```
    DB_HOST="postgres_db"
    DB_PORT=5432

    POSTGRES_PASSWORD="secret"
    POSTGRES_USER="user"
    POSTGRES_DB="animals"
```

### Run docker compose

```
docker-compose up -d
```

## Author

@steelWinds

## LICENSE

See license in LICENSE file from root dir
