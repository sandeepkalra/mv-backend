package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
)

func (am *AuthModule) ChangeOldPassword(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ChangeOldPasswordReq{Email: "", DigitLock: 0, OldPassword: "", NewPassword: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		len(request.OldPassword) == 0 ||
		len(request.NewPassword) == 0 ||
		request.DigitLock == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	person, err := models.People(am.DataBase, qm.Where("email=? AND password=? AND digit_lock = ?", request.Email, request.OldPassword, request.DigitLock)).One()
	if err == nil && person == nil {
		out.Msg = " entry credential failed  "
		return
	}

	if *(person.IsBlocked.Ptr()) == 1 {
		out.Msg = " Signup is still incomplete "
		return
	}

	person.Password = null.StringFrom(request.NewPassword)

	if err := person.Update(am.DataBase); err != nil {
		out.Msg = err.Error()
		return
	}

	out.Msg = "ok"
	out.Code = 0
	return
}
