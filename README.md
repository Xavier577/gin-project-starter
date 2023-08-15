# Gin App

This is a project starter for [gin](https://gin-gonic.com) apps

# How to use this template
This you can use the [temaple button](https://github.com/new?template_name=gin-project-starter&template_owner=Xavier577) to create a new repo or you can use degit.

```shell
$ npx degit https://github.com/Xavier577/gin-project-starter.git <project-name>
```

## Running the Server

### Tidy dependencies

```shell
$ go mod tidy
```

### Starting the server

```shell
$ go run main.go
```

## Folder structure
```shell
.
├── LICENSE
├── README.md
├── config
│   └── env
│       └── env.go
├── dbconf.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── migrations
│   └── 20230809154912_dbinit.sql
├── pkg
│   └── number
│       └── number.go
└── users
    ├── handler.go
    └── route.go

```

### Folders and file to take note of

- [docs](./docs): generated by [swag](https://github.com/swaggo/swag) for api documentation
- [migrations](./migrations): contains migrations for the postgres ([Goose](https://bitbucket.org/liamstask/goose/src/master) was used to generate and run the migration files but the migrations are handwritten)
- [dbconf.yml](./dbconf.yml): config file for goose

## Generate swagger docs with [Swag](https://github.com/swaggo/swag)

To generate swagger docs for your APIs with swag, run:

```shell
$ swag init -g main.go
```

## Migrations with [Goose](https://bitbucket.org/liamstask/goose/src/master)

### Creating Migrations

You can create migrations in either go or sql. SQL is preferred in this project.
To generate migrations with in SQL with [goose](https://bitbucket.org/liamstask/goose/src/master), run:

```shell
$ goose -path . create <your-migration-name>
```

### Running up migrations

```shell
$ goose -path . up -env <environment> # (check dbconf.yml to see configured environments)
```

### Running down migrations

```shell
$ goose -path . down -env <environment> # (check dbconf.yml to see configured environments)
```

