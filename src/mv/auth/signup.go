package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
)

func (am *AuthModule) Signup(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := SignupReq{Email: "", Password: "", DigitLock: 0, FName: "", LName: "", IsBlocked: true}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		len(request.Password) == 0 ||
		len(request.LName) == 0 ||
		len(request.FName) == 0 ||
		request.DigitLock == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	count, err := models.People(am.DataBase, qm.Where("email=?", request.Email)).Count()
	if err == nil && count > 0 {
		out.Msg = " entry already exist "
		return
	}

	// CREATE NEW ENTRY
	person := models.Person{
		Email:        null.StringFrom(request.Email),
		FName:        null.StringFrom(request.FName),
		LName:        null.StringFrom(request.LName),
		DigitLock:    null.IntFrom(request.DigitLock),
		OneTimeToken: null.StringFrom(gocql.TimeUUID().String()),
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