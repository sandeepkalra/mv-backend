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

func (am *AuthModule) ChangeOldDigitLock(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ChangeOldDigitLockReq{Email: "", OldDigitLock: 0, Password: "", NewDigitLock: 0}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		request.OldDigitLock == 0 ||
		request.NewDigitLock == 0 ||
		len(request.Password) == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	person, err := models.People(am.DataBase, qm.Where("email=? AND password=? AND digit_lock = ?", request.Email, request.Password, request.OldDigitLock)).One()
	if err == nil && person == nil {
		out.Msg = " entry credential failed  "
		return
	}

	if *(person.IsBlocked.Ptr()) == 1 {
		out.Msg = " Signup is still incomplete "
		return
	}

	person.DigitLock = null.IntFrom(request.NewDigitLock)

	if err := person.Update(am.DataBase); err != nil {
		out.Msg = err.Error()
		return
	}

	out.Msg = "ok"
	out.Code = 0
	return
}
