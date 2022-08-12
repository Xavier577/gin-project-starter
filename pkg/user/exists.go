package user

import "gin-project-starter/pkg/repositories"

func AlreadyExistsWithEmail(email string) bool {
	return repositories.User.Exists("email = ?", email)
}
