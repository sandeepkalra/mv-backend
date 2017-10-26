package utils

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"hash"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = ""
	DB_DBNAME   = "mvdb"
	DB_TYPE     = "mysql"
)

func IsPhone(email_or_ph string) bool {
	email_or_ph = strings.Replace(email_or_ph, " ", "", -1)
	fmt.Println(email_or_ph)
	isPh := true
	for i := 0; i < len(email_or_ph); i++ {
		if email_or_ph[i] >= '0' && email_or_ph[i] <= '9' {
			continue
		} else {
			isPh = false
			break
		}
	}
	return isPh
}

func GetCryptPassword(h hash.Hash, password string) string {
	h.Reset()
	h.Write([]byte(password))
	s := string(h.Sum(nil))
	return s
}

func InitCrypt256() hash.Hash {
	return sha256.New()
}

func InitDB() (*sql.DB, error) {
	return sql.Open(DB_TYPE, DB_USER+":"+DB_PASSWORD+"@/"+DB_DBNAME)
}
