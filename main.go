package main

import (
	"database/sql"
	"log"
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
	selDB, err := db.Query("SELECT * FROM test.todoapp ORDER BY id DESC")
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

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
