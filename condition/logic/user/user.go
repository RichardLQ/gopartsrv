package user

import (
	"gopartsrv/condition/model"
)

func UserInfo(userId string) (*model.Users, error) {
	var user model.Users
	user.Id = userId
	list, err := user.Find()
	if err != nil {
		return &model.Users{}, err
	}
	return list, nil
}
