package utils

import (
	"net/http"
	"time"
)

// SetCookie sets server side user session cookie to http header using ResponseWriter.
func SetCookie(w http.ResponseWriter, sessionID string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "SecretSessID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie)
}

// GetCookie gets cookie from the http request.
func GetCookie(r *http.Request) (sessionID string, e error) {
	c, e := r.Cookie("SecretSessID")
	if e != nil {
		return "", e
	}
	return c.Value, nil
}
