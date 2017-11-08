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

//ChangeOldDigitLock changes old digit_lock with new.
func (am *AuthModule) ChangeOldDigitLock(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ChangeOldDigitLockReq{Email: "", OldDigitLock: "", NewDigitLock: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		len(request.OldDigitLock) == 0 ||
		len(request.NewDigitLock) == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	person, err := models.People(am.DataBase, qm.Where("email=?", request.Email)).One()
	if err != nil || person == nil {
		out.Msg = " entry credential failed  "
		return
	}

	if person.IsBlocked.Int8 == 1 {
		out.Msg = " Signup is still incomplete "
		return
	}

	if b, e := utils.CheckPasswordHashes(request.OldDigitLock, person.DigitLock.String); b != true {
		out.Msg = " digit, password do not match ; " + e.Error()
		return
	}

	person.DigitLock = null.StringFrom(utils.GetCryptPassword(request.NewDigitLock))

	if err := person.Update(am.DataBase); err != nil {
		out.Msg = "failed" + err.Error()
		return
	}

	out.Msg = "ok"
	out.Code = 0
	return
}
