package user_service

import (
	"log"
	"tutorial/model"
	"tutorial/config/database"
)

type DBUserService struct {}

func (DBUserService) FindUser(User *user_model.Users, id int) error  {
	var err error
	err = database.DB.Where("id = ?", id).First(User).Error
	log.Print(err)
	if err != nil {
		log.Print("DEBUG: DB find error  ", err)
		return err
	}
	return nil
}