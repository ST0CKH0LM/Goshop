package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tem *template.Template
)

func connect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/goshop")
	checkerr(err)
	return db
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("=====================")
	fmt.Println("|| Program now run ||")
	fmt.Println("=====================")
	var err error
	tem, err = template.ParseGlob("template/*")
	checkerr(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/showdata", showdata)
	http.HandleFunc("/adddata", adddata)
	http.HandleFunc("/sortshow", sortshow)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tem.ExecuteTemplate(w, "index.html", nil)
}

type product struct {
	Id, Price                  int
	Name, Gen, Col, Patt, Size string
}

func showdata(w http.ResponseWriter, r *http.Request) {
	db := connect()
	rows, err := db.Query("SELECT * FROM db_category")
	checkerr(err)
	var sel []product
	for rows.Next() {
		var s product
		err = rows.Scan(&s.Id, &s.Name, &s.Gen, &s.Col, &s.Patt, &s.Size, &s.Price)
		checkerr(err)
		sel = append(sel, s)
	}
	tem.ExecuteTemplate(w, "index.html", sel)
}

func sortshow(w http.ResponseWriter, r *http.Request) {
	db := connect()
	showonly := r.FormValue("showonly")
	switch {
	case showonly == "10":
		rows, err := db.Query("SELECT * FROM db_category limit ?", showonly)
		checkerr(err)
		var sel []product
		for rows.Next() {
			var s product
			err = rows.Scan(&s.Id, &s.Name, &s.Gen, &s.Col, &s.Patt, &s.Size, &s.Price)
			checkerr(err)
			sel = append(sel, s)
		}
		tem.ExecuteTemplate(w, "index.html", sel)
	case showonly == "20":
		rows, err := db.Query("SELECT * FROM db_category limit ?", showonly)
		checkerr(err)
		var sel []product
		for rows.Next() {
			var s product
			err = rows.Scan(&s.Id, &s.Name, &s.Gen, &s.Col, &s.Patt, &s.Size, &s.Price)
			checkerr(err)
			sel = append(sel, s)
		}
		tem.ExecuteTemplate(w, "index.html", sel)
	case showonly == "50":
		rows, err := db.Query("SELECT * FROM db_category limit ?", showonly)
		checkerr(err)
		var sel []product
		for rows.Next() {
			var s product
			err = rows.Scan(&s.Id, &s.Name, &s.Gen, &s.Col, &s.Patt, &s.Size, &s.Price)
			checkerr(err)
			sel = append(sel, s)
		}
		tem.ExecuteTemplate(w, "index.html", sel)
	}
}

func adddata(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		tem.ExecuteTemplate(w, "add.html", nil)
		return
	}
	if r.Method == "POST" {
		db := connect()
		name := r.FormValue("name")
		price := r.FormValue("price")
		gender := r.FormValue("gender")
		size := r.FormValue("size")
		color := r.FormValue("color")
		pattern := r.FormValue("pattern")
		var err error
		_, err = db.Exec("INSERT INTO db_category(name,price,gender,size,color,pattern) VALUES(? , ? , ? , ? , ? , ?)", name, price, gender, size, color, pattern)
		checkerr(err)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
