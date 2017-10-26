package utils

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, sessionID string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "SecretSessID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request) (sessionID string, e error) {
	if c, e := r.Cookie("SecretSessID"); e != nil {
		return "", e
	} else {
		return c.Value, nil
	}
}
