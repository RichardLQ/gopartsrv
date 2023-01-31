package user

import (
	"gopartsrv/condition/model"
)

func UserInfo(userId,openid string) (*model.Users, error) {
	var user model.Users
	user.Id = userId
	user.Openid = openid
	list, err := user.Find()
	if err != nil {
		return &model.Users{}, err
	}
	return list, nil
}
