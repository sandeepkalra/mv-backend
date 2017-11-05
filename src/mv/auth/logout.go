package main

import (
	"mv/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (am *AuthModule) Logout(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	out := utils.GetResponseObject()

	out.Msg = "ok"
	out.Code = 0

	client_cookie, err := utils.GetCookie(req)
	if err == nil {
		/* we found the client cookie*/
		b, session_cookie := am.RedisDB.Get("ClientCookie", client_cookie)
		if b == true {
			/* For future identifier of session */
			_, email := am.RedisDB.Get("SessionCookieToEmail", session_cookie)
			am.RedisDB.Del("PersonName", email)
			am.RedisDB.Del("SessionCookie", email)
			am.RedisDB.Del("PersonId", email)
		}
	}

	out.Send(res)
	return
}
