## Requirements
- Golang
- gin-gonic
- PostgreSQL
- Docker

## How to run this

### 1. clone this repo
```bash
git clone https://github.com/eva81829/go-postgres.git
```

### 2. install dependencies
```bash
go mod init <project-name>
```

### 3. prepare the image
```bash
docker build -t <docker-account>/<project-name> .
docker pull postgres
```

### 4. run the DB container and create DB
```bash
docker run -d --name <db-container-name> -p 3306:5432 -e POSTGRES_PASSWORD=<password> postgres
docker exec -it <db-container-name> psql -U postgres -c "create role <db-owner> with login password '<password>';"
docker exec -it <db-container-name> psql -U postgres -c "create database <db-name> owner <db-owner>"
docker exec -it <db-container-name> psql -U postgres -c "\l"
docker exec -it <db-container-name> psql -U <db-owner> <db-name>

CREATE TABLE users (
   userid SERIAL PRIMARY KEY,
   username text NOT NULL
);
\q
```

### 4. run the Web API container
```bash
docker run --name <web-api-container-name> -p 8080:8080 --link <db-container-name> <docker-account>/<project-name>
```

### 4. check server is running
```bash
GET    http://localhost:8080/user
GET    http://localhost:8080/user/user_id
POSE   http://localhost:8080/user
PUT    http://localhost:8080/user/user_id
DELETE http://localhost:8080/user/user_id
```
