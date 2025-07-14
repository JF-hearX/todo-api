# todo-api

This is a Golang Todo API running on mysql

# Requirements

- Goland 1.24.5
- Docker and Docker compose

# Running just database

clone example.env to .env

```
cp example.env .env
```

```
docker compose up -d --build todo-api-db   
```

# Edit config.yam

Clone the config.yml.sample to config.yml

```
cp config.yml.sample config.yml
```

Edit the dsn (Data Source Name) and the server port.
Match the DSN to the .env.
should look like the following example:
`todo:secret@tcp(127.0.0.1:3306)/todo?multiStatements=true&parseTime=true`

Replace the 