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

func InitServer() (*ItemModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb")
	if err != nil {
		return nil, err
	}
	b, redis := utils.FastMemInit("127.0.0.1")
	if b != true {
		db.Close()
		return nil, fmt.Errorf("failed to init redis")
	}
	return &ItemModule{DataBase: db, RedisDB: redis}, nil
}

func (im *ItemModule) ServerClose() {
	im.DataBase.Close()
	im.RedisDB.R.Close()
}

func (im *ItemModule) Middleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if strings.Contains(req.URL.Path, "/item/") != true {

	} else {

	}
	next(res, req)
	return
}

func (im *ItemModule) Handler() http.Handler {

	n := negroni.Classic()
	//n.Use(c)
	n.Use(negroni.HandlerFunc(im.Middleware))
	r := httprouter.New()

	/*********** AUTH *************/
	r.POST("/item/add_item", im.AddItem)

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
	listen_str := "0.0.0.0:9502"
	fmt.Println("ITEM MICROSERVICES LISTENING AT AT ", listen_str)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listen_str, srv.Handler())
}
