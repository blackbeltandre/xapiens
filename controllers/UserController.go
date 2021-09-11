package controllers

import (
	"encoding/json"
	"xapiens/models"
	u "xapiens/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateUsers = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	users := &models.Users{}

	err := json.NewDecoder(r.Body).Decode(users)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	users.UserID = user
	resp := users.Create()
	u.Respond(w, resp)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetUsers(id)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetUsersDetail = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.GetUsersDetail(id,UserID)
	resp := u.Message(true, "Display Detail User Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var DeleteUsers = func(w http.ResponseWriter, r *http.Request) {
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := models.DeleteUsers(id,UserID)
	resp := u.Message(true, "Delete User Success")
	resp["data"] = data
	u.Respond(w, resp)
}


var UpdateUsers = func(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	user := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request User body"))
		return
	}
	idnya := r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idnya)
	fmt.Println(id)
	UserID := r.Context().Value("user").(uint)
	data := user.UpdateUsers(id,UserID)
	resp := u.Message(true, "Success Update User")
	resp["data"] = data
	u.Respond(w, resp)
}


