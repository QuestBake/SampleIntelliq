package service

import (
	"pracSpace/restHandler_Gin/app/model"
	"pracSpace/restHandler_Gin/app/repo"
)

//FindAllUsers address model
func FindAllUsers() (model.Users, error) {
	userRepo := repo.NewUserRepository()
	users, err := userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//AuthenticateUser login
func AuthenticateUser(user *model.User) (*model.User, error) {
	userRepo := repo.NewUserRepository()
	loggedUser, err := userRepo.FindByNameAndMobile(user)
	if err != nil {
		return nil, err
	}
	return loggedUser, nil
}

//SearchUsersWithSearchIndex search
func SearchUsersWithSearchIndex(term string) (model.Users, error) {
	userRepo := repo.NewUserRepository()
	users, err := userRepo.SearchByIndex(term)
	if err != nil {
		return nil, err
	}
	return users, nil
}
