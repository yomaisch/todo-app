package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Todo struct {
	Id   int
	Task string
}

func dbConn() (db *sql.DB) {
	psqlInfo := "host=localhost user=YoshimasaIshino password=yonce dbname=yonce port=5432 sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*")) //ParseGlobは、パターンによってマッチしたファイルのリストを持つParseFilesを呼び出すことと同じ

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM test.todo;")
	if err != nil {
		panic(err.Error())
	}
	td := Todo{}
	res := []Todo{}
	for selDB.Next() {
		var id int
		var task string
		err = selDB.Scan(&id, &task)
		if err != nil {
			panic(err.Error())
		}
		td.Id = id
		td.Task = task
		res = append(res, td)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		crtForm, err := db.Prepare("INSERT INTO test.todo(task) VALUES($1);")
		if err != nil {
			panic(err.Error())
		}
		crtForm.Exec(task)
	}
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/new", New)
	http.HandleFunc("/create", Create)
	http.ListenAndServe("localhost:8080", nil)
}
