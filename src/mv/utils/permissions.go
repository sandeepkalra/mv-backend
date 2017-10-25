

package utils

import (
	"net/http"
	"strconv"
)

func IsAllowed(req *http.Request) (email, user_type string, u_id int64, allowed bool) {
	email = req.Header.Get(AUTH_USER_EMAIL)
	user_type = req.Header.Get(AUTH_USER_TYPE)
	usrIdStr := req.Header.Get(AUTH_USER_USERID)
	u_id, _ = strconv.ParseInt(usrIdStr, 10, 64)
	isUserBlocked := (req.Header.Get(AUTH_USER_BLOCKED) != "0")
	allowed = (isUserBlocked == false)
	return
}
