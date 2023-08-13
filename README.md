# Cake Store

A CRUD API project for your favorite cakes. Using redis + docker

## How to start

1. Pull the repo

```bash
git pull https://github.com/bagasss3/cake-store
```

2. Start the go container and mysql container

```bash
docker-composer up  -d
```

3. Create database for cake store

```bash
docker exec -it godockerDB mysql -u root -p
```

```bash
CREATE DATABASE cakestore
```

4. Migrate the sql script from local project - change config.yml

```bash
database:
  host: "db:3306"

//to

database:
  host: "localhost:3306"
```

5. Migrate the sql script from local project - run migrate cmd

```bash
go run main.go migrate --direction=up
```

6. The APIs should be ready to use!

## Command Usage

```bash
# start server
go run main.go server

# start migrate sql scripts
go run main.go migrate

```
