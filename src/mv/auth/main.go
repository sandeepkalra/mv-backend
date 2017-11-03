package main

import (
	"database/sql"
	"fmt"
	_ "mv/models"

	"github.com/Codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"mv/utils"
	"net/http"
	"strings"

	_ "github.com/volatiletech/sqlboiler/queries/qm"
	_ "gopkg.in/volatiletech/null.v6"
)

func InitServer() (*AuthModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb")
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

func (am *AuthModule) ServerClose() {
	am.DataBase.Close()
	am.RedisDB.R.Close()
}

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

func (am *AuthModule) Handler() http.Handler {

	n := negroni.Classic()
	//n.Use(c)
	n.Use(negroni.HandlerFunc(am.Middleware))
	r := httprouter.New()

	/*********** AUTH *************/

	r.POST("/auth/signup", am.Signup)
	r.POST("/auth/validate_token", am.ValidateSignupToken)
	r.POST("/auth/change_old_password", am.ChangeOldPassword)
	r.POST("/auth/change_old_fourdigitlock", am.ChangeOldDigitLock)
	r.POST("/auth/login", am.Login)
	r.POST("/auth/logout", am.Logout)
	//r.POST("/auth/forgot_password", am.ForgotPassword)

	/* ********* LAUNCH ************/
	n.UseHandler(r)
	return n
}

func main() {
	srv, e := InitServer()
	if e != nil {
		return
	}

	defer srv.ServerClose()

	/************* PREPARE TO LAUNCH *************/
	listen_str := "0.0.0.0:9501"
	fmt.Println("AUTH MICROSERVICES LISTENING AT AT ", listen_str)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listen_str, srv.Handler())
}
