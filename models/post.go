package models

import (
	"fmt"
	u "go-blog/utils"

	"github.com/jinzhu/gorm"
)

// Post estructura de la tabla pensamientos
type Post struct {
	gorm.Model
	Title  string
	Body   string
}

// Create funcion para crear un post
func (post *Post) Create() map[string]interface{} {

	GetDB().Debug().Create(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

// Update funcion para actualizar un post
func (post *Post) Update() map[string]interface{} {

	GetDB().Save(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

// Delete funcion para eliminar un post
func (post *Post) Delete() map[string]interface{} {

	GetDB().Delete(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

// GetPost funcion para obtener un post
func GetPost(id string) *Post {

	post := &Post{}
	err := GetDB().First(&post, id).Error
	if err != nil {
		return nil
	}
	return post
}

// GetPosts funcion para obtener todos los post
func GetPosts() []*Post {

	post := make([]*Post, 0)
	err := GetDB().Find(&post).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return post
}
