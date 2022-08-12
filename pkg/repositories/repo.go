package repositories

import "gin-project-starter/pkg/repositories/orm"

var User IUser = &user{DbClient: orm.DbClient}
