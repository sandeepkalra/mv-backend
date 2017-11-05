package utils

import (
	"database/sql"
	"fmt"
	"strings"


	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func GetCryptPassword(password string) string {
	pasBytes, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		 //TODO: Properly handle error
		log.Fatal(e)
	}
	return string(pasBytes)
}

func CheckPasswordHashes(request_password, db_password string) (bool, error) {
	if e:= bcrypt.CompareHashAndPassword([]byte(db_password),[]byte(request_password)) ; e != nil {
		return false, e
	}
	return true, nil
}

func InitDB() (*sql.DB, error) {
	return sql.Open(DB_TYPE, DB_USER+":"+DB_PASSWORD+"@/"+DB_DBNAME)
}
