package main

import (

	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	//"github.com/bmizerany/pat"
	"log"
	_ "github.com/go-sql-driver/mysql"
	sql "database/sql"
)

type FileData struct {
	ID int
	UID int
	Title string
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
	CheckErr(err)

	// query
	rows, err := db.Query("SELECT id, uinfo_id, title FROM sharedata limit 0, 100")
	CheckErr(err)

	fds := []FileData{}

	for rows.Next() {
		var id int
		var uid int
		var title string
		err = rows.Scan(&id, &uid, &title)
		CheckErr(err)
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
	db, err := sql.Open("mysql", "root@/baidu?charset=utf8")
	CheckErr(err)

	// query
	rows, err := db.Query("SELECT * FROM uinfo")
	CheckErr(err)

	for rows.Next() {
		var id int
		var uk int
		var username string
		var avatarurl string
		err = rows.Scan(&id, &uk, &username, &avatarurl)
		CheckErr(err)
		fmt.Println(id)
		fmt.Println(uk)
		fmt.Println(username)
		fmt.Println(avatarurl)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Println("Error...", err)
	}
}

func main() {
	DBTest()


	mx := mux.NewRouter()

	mx.HandleFunc("/", Index)
	mx.HandleFunc("/list/{type}/{page:[0-1000]+}", ListFile)
	mx.HandleFunc("/search/{type}/{keyword}", SearchFile)
	mx.HandleFunc("/file/{id}", ShowFile)
	mx.HandleFunc("/user/{id}/{page}", ShowUser)
	mx.PathPrefix("/static").Handler(http.FileServer(http.Dir("./")))

	log.Println("Listening...")
	http.ListenAndServe(":8080", mx)

}
