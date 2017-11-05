package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (am *AuthModule) Login(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := LoginReq{Email: "", Password: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.Email) == 0 ||
		len(request.Password) == 0 {
		out.Msg = " Empty fields not allowed "
		return
	}

	person, err := models.People(am.DataBase, qm.Where("email=?", request.Email)).One()
	if err != nil || person == nil {
		out.Msg = " entry does not  exist "
		return
	}

	if b, e := utils.CheckPasswordHashes(request.Password, person.Password.String); b != true {
		out.Msg = e.Error()
		return
	}

	if person.IsBlocked.Int8 == 1 {
		out.Msg = " Sorry, you are blocked as of now. Please complete your signup, or, unlock yourself to login"
		return
	}

	name := person.FName.String + " " + person.LName.String

	out.Msg = "ok"
	out.Code = 0

	session_cookie := gocql.TimeUUID().String()
	client_cookie := gocql.TimeUUID().String()

	out.Response = map[string]interface{}{
		"client_cookie": client_cookie,
	}

	/* For future identifier of session */
	am.RedisDB.TimedAdd("PersonName", request.Email, name)
	am.RedisDB.TimedAdd("ClientCookie", client_cookie, session_cookie)
	am.RedisDB.TimedAdd("SessionCookieToEmail", session_cookie, request.Email)
	am.RedisDB.TimedAdd("SessionCookie", request.Email, session_cookie)
	am.RedisDB.TimedAdd("PersonId", request.Email, strconv.FormatInt(person.ID, 10))

	utils.SetCookie(res, client_cookie)

	return
}
