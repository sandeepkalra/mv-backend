package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
)

// ValidateSignupToken validate the signup one time token, thus completing the signup process.
func (am *AuthModule) ValidateSignupToken(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ValidateSingupReq{Email: "", Token: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		len(request.Token) == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	person, err := models.People(am.DataBase, qm.Where("email=? AND one_time_token=?", request.Email, request.Token)).One()
	if err != nil || person == nil {
		out.Msg = " entry does not  exist " + err.Error()
		fmt.Println(err.Error())
		return
	}

	person.IsBlocked = null.Int8From(0)
	person.OneTimeToken = null.StringFrom("")

	fmt.Println("reset ing blocked->unblocked, one-time token to null")

	if err := person.Update(am.DataBase); err != nil {
		out.Msg = err.Error()
		return
	}

	out.Msg = "ok"
	out.Code = 0
	return
}
