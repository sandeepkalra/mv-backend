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

	out.Send(res)
	return
}
