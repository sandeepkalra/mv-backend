package main

import (
	"../utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Logout lets user logout
func (am *AuthModule) Logout(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	out := utils.GetResponseObject()

	out.Msg = "ok"
	out.Code = 0

	clientCookie, err := utils.GetCookie(req)
	if err == nil {
		/* we found the client cookie*/
		b, sessionCookie := am.RedisDB.Get("ClientCookie", clientCookie)
		if b == true {
			/* For future identifier of session */
			_, email := am.RedisDB.Get("SessionCookieToEmail", sessionCookie)
			am.RedisDB.Del("PersonName", email)
			am.RedisDB.Del("SessionCookie", email)
			am.RedisDB.Del("PersonID", email)
		}
	}

	out.Send(res)
	return
}
