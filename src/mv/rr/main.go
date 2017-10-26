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

func InitServer() (*RRModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb")
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

func (rr *RRModule) ServerClose() {
	rr.DataBase.Close()
	rr.RedisDB.R.Close()
}

func (rr *RRModule) Middleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if strings.Contains(req.URL.Path, "/rr/") != true {

	} else {

	}
	next(res, req)
	return
}

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

func main() {
	srv, e := InitServer()
	if e != nil {
		return
	}

	defer srv.ServerClose()

	/************* PREPARE TO LAUNCH *************/
	listen_str := "0.0.0.0:9503"
	fmt.Println("RR MICROSERVICES LISTENING AT AT ", listen_str)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listen_str, srv.Handler())
}
