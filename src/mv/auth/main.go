package main

import (
	_ "../models"
	"database/sql"
	"fmt"

	"github.com/Codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"../utils"
	"net/http"
	"strings"

	_ "github.com/volatiletech/sqlboiler/queries/qm"
	_ "gopkg.in/volatiletech/null.v6"
)

// InitServer intialize auth module
func InitServer() (*AuthModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb?parseTime=true")
	if err != nil {
		return nil, err
	}
	b, redis := utils.FastMemInit("127.0.0.1")
	if b != true {
		db.Close()
		return nil, fmt.Errorf("failed to init redis")
	}

	return &AuthModule{DataBase: db, RedisDB: redis}, nil
}

// ServerClose closes the handles of the server, handle destroy of objects.
func (am *AuthModule) ServerClose() {
	am.DataBase.Close()
	am.RedisDB.R.Close()
}

// Middleware is middle-ware of the http router.
func (am *AuthModule) Middleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	/* CORS */
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if strings.Contains(req.URL.Path, "/auth/") != true {

	} else {

	}
	next(res, req)
	return
}

// Handler is the main handler of the http router.
func (am *AuthModule) Handler() http.Handler {

	n := negroni.Classic()
	//n.Use(c)
	n.Use(negroni.HandlerFunc(am.Middleware))
	r := httprouter.New()

	/*********** AUTH *************/

	r.POST("/auth/signup", am.Signup)
	r.POST("/auth/validate_token", am.ValidateSignupToken)
	r.POST("/auth/change_old_password", am.ChangeOldPassword)
	r.POST("/auth/change_old_digit_lock", am.ChangeOldDigitLock)
	r.POST("/auth/reset_password", am.ResetPassword)
	r.POST("/auth/reset_digit_lock", am.ResetDigitLock)
	r.POST("/auth/login", am.Login)
	r.POST("/auth/logout", am.Logout)

	/* ********* LAUNCH ************/
	n.UseHandler(r)
	return n
}

// main Auth control starts here.
func main() {
	srv, e := InitServer()
	if e != nil {
		return
	}

	defer srv.ServerClose()

	/************* PREPARE TO LAUNCH *************/
	listenStr := "0.0.0.0:9501"
	fmt.Println("AUTH MICROSERVICES LISTENING AT AT ", listenStr)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listenStr, srv.Handler())
}
