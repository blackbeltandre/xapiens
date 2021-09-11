package models

import (
	"github.com/dgrijalva/jwt-go"
    "gorm.io/gorm"
	u "xapiens/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	Username    string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {


	if len(account.Username) < 6 {
		return u.Message(false, "Username is required"), false
	}
	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Username must be unique
	temp := &Account{}

	//check for errors and duplicate Username
	// err := GetDB().Table("accounts").Where("Username = ?", account.Username).First(temp).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return u.Message(false, "Connection error. Please retry"), false
	// }
	if temp.Username != "" {
		return u.Message(false, "Username address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = account.Password //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(username, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("username = ?", username).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Username address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = account.Password

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Username == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
