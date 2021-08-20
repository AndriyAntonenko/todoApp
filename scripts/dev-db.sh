# /bin/sh

# Run container
docker run -p 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -d --name todo-app-dev-db postgres

# Run migrations
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up