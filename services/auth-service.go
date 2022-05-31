package services

import (
	"github.com/gestaoFrota/model"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	FindByEmail(email string) model.Users
	IsDuplicateEmail(email string) bool
}

func FindByEmail(email string) model.Users{
	var user model.Users
	OpenDatabase()
	Db.Find(&user, "email = ?", email).Take(&user)
	return user
}

func IsDuplicateEmail(email string) bool {
	var user model.Users
	OpenDatabase()
	return Db.Find(&user, "email = ?", email).Take(&user).Error == nil
}

