package utils

import (
	"database/sql"
	"fmt"
	"strings"

	"log"

	"golang.org/x/crypto/bcrypt"
)

// enums for the DB connection
const (
	DBUser       = "root"
	DBPassword   = ""
	DBSchemaName = "mvdb"
	DBType       = "mysql"
)

// IsPhone is to find if this is phone number or email.
func IsPhone(emailOrPhone string) bool {
	emailOrPhone = strings.Replace(emailOrPhone, " ", "", -1)
	fmt.Println(emailOrPhone)
	isPh := true
	for i := 0; i < len(emailOrPhone); i++ {
		if emailOrPhone[i] >= '0' && emailOrPhone[i] <= '9' {
			continue
		} else {
			isPh = false
			break
		}
	}
	return isPh
}

//GetCryptPassword creates bcrypt password hash
func GetCryptPassword(password string) string {
	pasBytes, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		//TODO: Properly handle error
		log.Fatal(e)
	}
	return string(pasBytes)
}

//CheckPasswordHashes check if given password belongs to same bcrypt hash or not
func CheckPasswordHashes(userGivenPassword, userDBPassword string) (bool, error) {
	if e := bcrypt.CompareHashAndPassword([]byte(userDBPassword), []byte(userGivenPassword)); e != nil {
		return false, e
	}
	return true, nil
}

//InitDB initializes handle for DBs
func InitDB() (*sql.DB, error) {
	return sql.Open(DBType, DBUser+":"+DBPassword+"@/"+DBSchemaName)
}
