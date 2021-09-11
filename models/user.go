package models

import (
	"fmt"
	"gorm.io/gorm"
	u "xapiens/utils"
)

type Users struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	Email   string `json:"email"`
	Password   string `json:"password"`
	Address  string `json:"address"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (user *Users) Validate() (map[string]interface{}, bool) {
	if user.Email  == "" {
		return u.Message(false, "Email  is required"), false
	}
	if len(user.Password) < 1 {
		return u.Message(false, "Password is required"), false
	}
	if len(user.Address) < 1 {
		return u.Message(false, "Address is required"), false
	}


	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (user *Users) Create() (map[string]interface{}) {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetUsernya(id uint) (*Users) {

	user := &Users{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUsers(UserID uint) ([]*Users) {

	user := make([]*Users, 0)
	err := GetDB().Table("users").Where("user_id = ?", UserID).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func GetUsersDetail(id int, UserID uint) ([]*Users) {
	user := make([]*Users, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}
func DeleteUsers(id int, UserID uint) ([]*Users) {
	user := make([]*Users, 0)
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Delete(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}
func (user *Users) UpdateUsers(id int, UserID uint) (map[string]interface{}) {

	if resp, ok := user.Validate(); !ok {
		return resp
	}
	err := GetDB().Where("id = ? and user_id = ?", id, UserID).Updates(Users{Email: user.Email,UserID: user.UserID,Password: user.Password,Address: user.Address}).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "Update User success")
	resp["user"] = user
	return resp
}