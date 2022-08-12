# Using this Template

To scaffold your project using this template.Click the use template button to create a new repo based on this template or run:

```
npx degit https://github.com/Xavier577/gin-project-starter.git <your-project-directory>
```

# Things to note while using this template

### Folder structure

Everything is structured to losely follow the hexagonal structure. You can edit to fit your needs. Everything is loosely coupled so you can swap out anything as need (swag gin for fiber or chi, swag gorm for xorm etc.)

```
.
├── LICENSE
├── README.md
├── cmd
│   └── main.go
├── database
│   └── connection.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yamls
├── go.mod
├── go.sum
├── pg
│   ├── dbconf.yml
│   └── migrations
│       └── 20220810124355_init.sql
└── pkg
    ├── config
    │   └── config.go
    ├── dtos
    │   └── newuser.go
    ├── env
    │   └── env.go
    ├── errors
    │   └── errors.go
    ├── handlers
    │   ├── auth
    │   │   └── auth.go
    │   ├── index
    │   │   └── index.go
    │   └── user
    │       └── create.go
    ├── middlewares
    │   └── auth.go
    ├── models
    │   └── users.go
    ├── repositories
    │   ├── orm
    │   │   └── client.go
    │   ├── repo.go
    │   └── user.go
    ├── routes
    │   └── routes.go
    └── user
        ├── create.go
        └── exists.go
```

### Files to take note of

- pg -> contains migrations for the postgres ([Goose](https://bitbucket.org/liamstask/goose/src/master) was used to generate and run the migration file but the migrations are hand written)
- docs -> generated by [swag](https://github.com/swaggo/swag) for api documentation
