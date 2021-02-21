package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "a", Email: "a@b.c", Age: 10}
	user2 := User{Name: "d", Email: "e@f.g", Age: 101}
	users := []User{user, user2}

	tmpl, err := template.New("Tmpl1").ParseFiles("./templates/tmpl1.tmpl", "./templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
