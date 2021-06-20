package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Todo struct {
	Id   int
	Task string
}

// func init() {
// 	connectdb()
// }

// func connectdb() (db *sql.DB) {
// 	psqlInfo := "host=localhost user=YoshimasaIshino password=yonce dbname=yonce port=5432 sslmode=disable"
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

func main() {
	// テンプレートの生成
	http.HandleFunc("/", index)
	http.HandleFunc("/", create)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// 試しにDBにサンプルデータを入れ込む
	t := Todo{Task: "second ch"}
	t.Create()
}

func (t *Todo) create((w http.ResponseWriter, r *http.Request)) (err error) {
	psqlInfo := "host=localhost user=YoshimasaIshino password=yonce dbname=yonce port=5432 sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles("templates/tmpl.html")
	t.Execute(w, )

	statement := "INSERT INTO test.todoapp(task) VALUES($1) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(t.Task).Scan(&t.Id)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/tmpl.html")
	t.Execute(w, "hello go")
}
