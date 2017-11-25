package utils

import (
	"net/http"
	"strconv"
)

// IsAllowed is part of middleware to http router.
func IsAllowed(req *http.Request) (email, userType string, uID int64, allowed bool) {
	email = req.Header.Get(AuthUserEmail)
	userType = req.Header.Get(AuthUserType)
	usrIDStr := req.Header.Get(AuthUserID)
	uID, _ = strconv.ParseInt(usrIDStr, 10, 64)
	isUserBlocked := (req.Header.Get(AuthIsUserBlocked) != "0")
	allowed = (isUserBlocked == false)
	return
}
