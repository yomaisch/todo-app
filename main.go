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

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)
	http.ListenAndServe("localhost:8080", nil)
}

func dbConn() (db *sql.DB) {
	psqlInfo := "host=localhost user=YoshimasaIshino password=yonce dbname=yonce port=5432 sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//ParseGlobは、パターンによってマッチしたファイルのリストを持つParseFilesを呼び出すことと同義
var tmpl = template.Must(template.ParseGlob("templates/*"))

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

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	uId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM test.todo WHERE id=$1", uId)
	if err != nil {
		panic(err.Error())
	}

	td := Todo{}
	for selDB.Next() {
		var id int
		var task string
		err = selDB.Scan(&id, &task)
		if err != nil {
			panic(err.Error())
		}
		td.Id = id
		td.Task = task
	}
	tmpl.ExecuteTemplate(w, "Edit", td)
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

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	td := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM test.todo WHERE id=$1;")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(td)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		udtForm, err := db.Prepare("UPDATE test.todo SET task=$1 WHERE id=$2;")
		if err != nil {
			panic(err.Error())
		}
		udtForm.Exec(task)
	}
	http.Redirect(w, r, "/", 301)
}
