package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"fmt"
	"github.com/gocql/gocql"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

// Signup signup a new user
func (am *AuthModule) Signup(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := SignupReq{Email: "", Password: "", DigitLock: "", FName: "", LName: "", IsBlocked: true}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = "signup: failed to decode incoming msg " + e.Error()
		out.Response = e.Error()
		return
	}

	if len(request.Email) == 0 ||
		len(request.Password) == 0 ||
		len(request.LName) == 0 ||
		len(request.FName) == 0 ||
		len(request.DigitLock) == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	count, err := models.People(am.DataBase, qm.Where("email=?", request.Email)).Count()
	if err == nil && count > 0 {
		out.Msg = " entry already exist "
		return
	}

	encrptedPassword := utils.GetCryptPassword(request.Password)
	encruptedDigitLock := utils.GetCryptPassword(request.DigitLock)
	fmt.Println("encrypted password ", encrptedPassword)
	// CREATE NEW ENTRY
	person := models.Person{
		Email:        null.StringFrom(request.Email),
		FName:        null.StringFrom(request.FName),
		LName:        null.StringFrom(request.LName),
		DigitLock:    null.StringFrom(encruptedDigitLock),
		Password:     null.StringFrom(encrptedPassword),
		CreatedOn:    null.TimeFrom(time.Now()),
		OneTimeToken: null.StringFrom(gocql.TimeUUID().String()),
		IsBlocked:    null.Int8From(1),
	}

	if err := person.Insert(am.DataBase); err != nil {
		out.Msg = err.Error()
		return
	}

	out.Msg = "ok"
	out.Response = person.OneTimeToken.Ptr()
	out.Code = 0
	return
}
