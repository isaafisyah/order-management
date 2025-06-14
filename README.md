# Record Transaction API

App Order Management, dibuat dengan Golang.

# docker

- docker-compose build
- docker-compose up -d
- docker logs multifinance
- docker exec -ti multifinance bash
- docker-compose stop (jika ingin stop docker)

# migration

- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- export PATH=$PATH:~/go/bin
- migrate create -ext sql -dir app/database/migrations create_table_tableName
- docker-compose run migrate
- docker-compose run migrate-down (untuk rollback)
