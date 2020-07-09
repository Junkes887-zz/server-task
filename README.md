# Server Task
Manager tasks

### Init
go-server `docker-compose up -d`

Front application: https://github.com/Junkes887/task-manager

### Migrations
Build image: `docker build -t migrator ./migrator/`

Run migrations: `docker run --network host migrator -path=/migrations/ -database "postgresql://{username}:{password}@{host}:{port}/postgres?sslmode=disable" up`
