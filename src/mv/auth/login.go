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

	person, err := models.People(am.DataBase, qm.Where("email=? AND password=?", request.Email, request.Password)).One()
	if err != nil || person == nil {
		out.Msg = " entry does not  exist "
		return
	}

	out.Msg = "ok"
	out.Code = 0
	cookie := gocql.TimeUUID().String()
	out.Response = map[string]interface{}{
		"cookie": cookie,
	}

	am.RedisDB.TimedAdd("SessionCookie", request.Email, cookie)
	am.RedisDB.TimedAdd("PersonId", request.Email, strconv.FormatInt(person.ID, 64))
	return
}
