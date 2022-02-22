package user_service

import (
	"errors"
	"tutorial/model"
)

type UserService struct {}

func (UserService) GetUser(id int) (user_model.UserModel, error) {
	var dbUserList [5]user_model.UserModel = [5]user_model.UserModel{
		{Uid: 1, Age: 10, Name: "test1"},
		{Uid: 2, Age: 20, Name: "test2"},
		{Uid: 3, Age: 30, Name: "test3"},
		{Uid: 4, Age: 40, Name: "test4"},
		{Uid: 5, Age: 50, Name: "test5"},
	}
	var selectedUser user_model.UserModel
	for _, value := range dbUserList {
		if value.Uid == id {
			selectedUser = value
			break
		}
	}
	if selectedUser.Uid == 0 {
		return selectedUser, errors.New("userが存在しません")
	}
	return selectedUser, nil
}