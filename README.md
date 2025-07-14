# Todo-Api

This is a Golang Todo API running on mysql (using mariadb image)

# Requirements

- Goland 1.24.5
- Docker and Docker compose

# Clone repository

using ssh`git@github.com:JF-hearX/todo-api.git` or https` https://github.com/JF-hearX/todo-api.git`

# Running just database

clone example.env to .env

```
cp example.env .env
```

update environments variable for .env file

```
docker compose up -d --build todo-api-db   
```

# Edit config.yml

Clone the config.yml.sample to config.yml

```
cp example.config.yml config.yml
```

Edit the dsn (Data Source Name) and the server port.
Match the DSN to the .env.
should look like the following example:
`todo:secret@tcp(127.0.0.1:3306)/todo?multiStatements=true&parseTime=true`

Replace the values matching .env file.

# Running the API in terminal
With Go 1.24.5 installed in the system

run in terminal`go run main.go migrate` to migrate the database

running the api in terminal `go run main.go api` to start the api

# Running Docker stack

If the database is migrated then run `docker compose up -d` to start the database and api

if the database is not migraded.  Can be migraged by using the terminal method above or change api to migrate in the compose.yaml

```
todo-api:
    build: .
    env_file:
      - ./.env
    # working_dir: /go/src/app
    command: api # change this to migrate
    container_name: todo-api
```

The container will run and stop.  Change back to `api` and run `docker compose up -d` to bring up the stack with the api and database running.



# Notes

The Commands that the Application can accept:

Can be viewed when running `go run main.go` or `go run main.go help`

```
Usage:
   [command]

Available Commands:
  api         Run http api service
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  migrate     Run http api service
  migratedown Run http api service

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```