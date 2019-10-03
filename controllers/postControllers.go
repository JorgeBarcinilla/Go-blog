package controllers

import (
	"go-blog/models"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	res := models.GetPosts()
	tmpl.ExecuteTemplate(w, "Index", res)
}

// Show ..
func Show(w http.ResponseWriter, r *http.Request) {
	nID := r.URL.Query().Get("id")

	emp := models.GetPost(nID)

	tmpl.ExecuteTemplate(w, "Show", emp)
}

// New ...
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit ...
func Edit(w http.ResponseWriter, r *http.Request) {
	nID := r.URL.Query().Get("id")

	emp := models.GetPost(nID)
	log.Println("EDIT: id: " + nID)
	tmpl.ExecuteTemplate(w, "Edit", emp)
}

//Insert ...
func Insert(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	post.Title = r.FormValue("title")
	post.Body = r.FormValue("body")

	post.Create()
	log.Println("INSERT: Title: " + post.Title + " | Body: " + post.Body)
	http.Redirect(w, r, "/", 301)
}

// Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	//post := &models.Post{}
	//ID, err := strconv.Atoi(r.FormValue("uid"))

	nID := r.URL.Query().Get("id")
	post := models.GetPost(nID)

	post.Title = r.FormValue("title")
	post.Body = r.FormValue("body")

	post.Update()

	log.Println("UPDATE: Title: " + post.Title + " | Body: " + post.Body)

	http.Redirect(w, r, "/", 301)
}

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	nID := r.URL.Query().Get("id")
	post := models.GetPost(nID)

	post.Delete()

	log.Println("DELETE: Title: " + post.Title + " | Body: " + post.Body)
	http.Redirect(w, r, "/", 301)
}
