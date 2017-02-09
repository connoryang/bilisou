package main

import (

	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
//	"log"
	_ "github.com/go-sql-driver/mysql"
	sql "database/sql"
//	"io/ioutil"
	"github.com/siddontang/go/log"
//	"regexp"
//	"encoding/json"
//	"time"
//	"github.com/garyburd/redigo/redis"
	"github.com/Unknwon/goconfig"
//	"strconv"
	"bytes"
//	"os"
//	"bufio"
//	"io"
//	"strings"
	m "model"
	u "utils"
)

type FileData struct {
	ID int
	UID int
	Title string
}


var db *sql.DB
var err error
var username, password, url, address, redis_Pwd, mode, logLevel, redis_db string
var redis_Database int
var ConfError error
var cfg *goconfig.ConfigFile

//Mysql Redis初始化
func Init() {
	cfg, ConfError = goconfig.LoadConfigFile("config.ini")
	if ConfError != nil {
		log.Error("配置文件config.ini不存在,请将配置文件复制到运行目录下")
	}
	logLevel, ConfError = cfg.GetValue("Log", "logLevel")
	if ConfError != nil {
		log.SetLevel(log.LevelInfo)
	} else {
		log.SetLevelByName(logLevel)
	}
	username, ConfError = cfg.GetValue("MySQL", "username")
	if ConfError != nil {
		log.Error("读取数据库username错误")
	}
	password, ConfError = cfg.GetValue("MySQL", "password")
	if ConfError != nil {
		log.Error("读取数据库password错误")
	}
	url, ConfError = cfg.GetValue("MySQL", "url")
	if ConfError != nil {
		log.Error("读取数据库url错误")
	}
	/*
	address, ConfError = cfg.GetValue("Redis", "address")
	if ConfError != nil {
		panic("读取数据库address错误")
	}
	redis_Pwd, ConfError = cfg.GetValue("Redis", "password")
	if ConfError != nil {
		panic("读取Redis password错误")
	}
	redis_db, ConfError = cfg.GetValue("Redis", "database")
	if ConfError != nil {
		redis_db = "0"
	}
	redis_Database, ConfError = strconv.Atoi(redis_db)
	if ConfError != nil {
		redis_Database = 0
	}
	*/
	var dataSourceName bytes.Buffer
	dataSourceName.WriteString(username)
	dataSourceName.WriteString(":")
	dataSourceName.WriteString(password)
	dataSourceName.WriteString("@")
	dataSourceName.WriteString(url)
	db, err = sql.Open("mysql", dataSourceName.String())
	if err != nil {
		log.Error(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Error Connection database...")
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
//	initRedisPool()
//	initWriteHasIndexKey();
}



func Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello, World!"))
	fds := GetFileData()
	render(w, "templates/index.html", fds)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func ListFile(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	t := vars["type"]
	p := vars["page"]
	w.Write([]byte(fmt.Sprintf("type is %s ", t)))
	w.Write([]byte(fmt.Sprintf("page is %s ", p)))
}


func SearchFile(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	t := vars["type"]
	kw := vars["keyword"]
	w.Write([]byte(fmt.Sprintf("type is %s ", t)))
	w.Write([]byte(fmt.Sprintf("keyword is %s ", kw)))
}

func ShowFile(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte(fmt.Sprintf("file id is %s ", id)))
}

func ShowUser(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	id := vars["id"]
	page := vars["page"]
	w.Write([]byte(fmt.Sprintf("user id is %s ", id)))
	w.Write([]byte(fmt.Sprintf("page id is %s ", page)))
}


func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetFileData() ([]FileData) {
	db, err := sql.Open("mysql", "root@/baidu?charset=utf8")
	u.CheckErr(err)

	// query
	rows, err := db.Query("SELECT id, uinfo_id, title FROM sharedata limit 0, 100")
	u.CheckErr(err)

	fds := []FileData{}

	for rows.Next() {
		var id int
		var uid int
		var title string
		err = rows.Scan(&id, &uid, &title)
		u.CheckErr(err)
//		fmt.Println(id)
//		fmt.Println(uid)
	//	fmt.Println(title)
		fd := FileData {
			ID : id,
			UID : uid,
			Title : title,
		}
		fds = append(fds, fd)
	}
	return fds
}


func DBTest() {

	// query
	rows, err := db.Query("SELECT id, uk, uname, avatar_url FROM uinfo")
	u.CheckErr(err)

	for rows.Next() {
		var id int
		var uk int
		var username string
		var avatarurl string
		err = rows.Scan(&id, &uk, &username, &avatarurl)
		u.CheckErr(err)
		fmt.Println(id)
		fmt.Println(uk)
		fmt.Println(username)
		fmt.Println(avatarurl)
	}
}

func main() {
	Init()
	//DBTest()
	m.GetShareVar(db, "4696533378915726131")

	mx := mux.NewRouter()

	mx.HandleFunc("/", Index)
	mx.HandleFunc("/list/{type}/{page:[0-1000]+}", ListFile)
	mx.HandleFunc("/search/{type}/{keyword}", SearchFile)
	mx.HandleFunc("/file/{id}", ShowFile)
	mx.HandleFunc("/user/{id}/{page}", ShowUser)
	mx.PathPrefix("/static").Handler(http.FileServer(http.Dir("./")))

	log.Info("Listening...")
	http.ListenAndServe(":8080", mx)

}
