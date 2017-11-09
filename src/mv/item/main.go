package main

import (
	"database/sql"
	"fmt"

	"mv/utils"
	"net/http"
	"strings"

	"github.com/Codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

// InitServer start item-server
func InitServer() (*ItemModule, error) {
	db, err := sql.Open("mysql", "root:@/mvdb?parseTime=true")
	if err != nil {
		fmt.Println("failed to connect db", err.Error())
		return nil, err
	}
	b, redis := utils.FastMemInit("127.0.0.1")
	if b != true {
		fmt.Println("failed to start fastmem")
		db.Close()
		return nil, fmt.Errorf("failed to init redis")
	}
	return &ItemModule{DataBase: db, RedisDB: redis}, nil
}

// ServerClose stops item server
func (im *ItemModule) ServerClose() {
	im.DataBase.Close()
	im.RedisDB.R.Close()
}

// Middleware middleware to http router.
func (im *ItemModule) Middleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	/* CORS */
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if strings.Contains(req.URL.Path, "/item/") != true {

	} else {

	}
	next(res, req)
	return
}

//Handler http handler
func (im *ItemModule) Handler() http.Handler {

	n := negroni.Classic()
	//n.Use(c)
	n.Use(negroni.HandlerFunc(im.Middleware))
	r := httprouter.New()

	/*********** AUTH *************/
	r.POST("/item/add_item", im.AddItem)
	r.POST("/item/update_item", im.UpdateItem)
	r.POST("/item/lookup_item", im.LookupItem)
	r.POST("/item/lookup_list", im.LookupList)

	/* ********* LAUNCH ************/
	n.UseHandler(r)
	return n
}

// main item-module
func main() {
	srv, e := InitServer()
	if e != nil {
		return
	}

	defer srv.ServerClose()

	/************* PREPARE TO LAUNCH *************/
	listenStr := "0.0.0.0:9502"
	fmt.Println("ITEM MICROSERVICES LISTENING AT AT ", listenStr)

	/************* SET THE BALL ROLLING  *************/
	http.ListenAndServe(listenStr, srv.Handler())
}
