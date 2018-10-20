package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "root"
	DBDbase = "cms"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	GUID       string
}

func (p *Page) ContentPreview() template.HTML {
	if len(p.Content) <= 30 {
		return p.Content
	}

	return p.Content[0:27] + "..."
}

func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)
	err := database.QueryRow("SELECT page_guid, page_title, page_content, page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.GUID, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Could not get page:", pageGUID)
	} else {
		thisPage.Content = template.HTML(thisPage.RawContent)
		template, err := template.ParseFiles("./templates/page.html")
		if err != nil {
			log.Println(err.Error())
		} else {
			template.Execute(w, thisPage)
		}
	}
}

func redirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	Pages := []Page{}
	pages, err := database.Query("SELECT page_guid, page_title, page_content, page_date FROM pages ORDER BY ? DESC", "page_date")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	defer pages.Close()
	for pages.Next() {
		thisPage := Page{}
		pages.Scan(&thisPage.GUID, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
		thisPage.Content = template.HTML(thisPage.RawContent)
		Pages = append(Pages, thisPage)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Pages)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error())

	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", servePage)
	routes.HandleFunc("/", redirIndex)
	routes.HandleFunc("/home", serveIndex)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
