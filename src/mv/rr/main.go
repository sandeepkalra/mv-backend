package main

import (
	"database/sql"
	"fmt"

	"github.com/Codegangsta/negroni"
	"github.com/julienschmidt/httprouter"

	"mv/utils"
	"net/http"
	"strings"
)

// InitServer intialize RR module
func InitServer() (*RRModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb?parseTime=true")
	if err != nil {
		return nil, err
	}
	b, redis := utils.FastMemInit("127.0.0.1")
	if b != true {
		db.Close()
		return nil, fmt.Errorf("failed to init redis")
	}
	return &RRModule{DataBase: db, RedisDB: redis}, nil
}

// ServerClose destroys RR module
func (rr *RRModule) ServerClose() {
	rr.DataBase.Close()
	rr.RedisDB.R.Close()
}

// Middleware middle-ware to http router.
func (rr *RRModule) Middleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	/* CORS */
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if strings.Contains(req.URL.Path, "/rr/") != true {

	} else {

	}
	next(res, req)
	return
}

// Handler handles http request, http router main handler
func (rr *RRModule) Handler() http.Handler {

	n := negroni.Classic()
	//n.Use(c)
	n.Use(negroni.HandlerFunc(rr.Middleware))
	r := httprouter.New()

	/*********** AUTH *************/
	r.POST("/rr/add_rr", rr.AddRR)
	r.POST("/rr/list_rr", rr.ListRR)
	/* ********* LAUNCH ************/
	n.UseHandler(r)
	return n
}

// main - main RR control starting point.
func main() {
	srv, e := InitServer()
	if e != nil {
		return
	}

	defer srv.ServerClose()

	/************* PREPARE TO LAUNCH *************/
	listenStr := "0.0.0.0:9503"
	fmt.Println("RR MICROSERVICES LISTENING AT AT ", listenStr)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listenStr, srv.Handler())
}
