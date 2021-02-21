# Template

틀에 input 을 넣어서 다양한 output 을 만들게 해줌

```go
package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{Name: "a", Email: "a@b.c", Age: 100}
	tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge : {{.Age}}")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, user)
}

```

template 파일을 생성한다음에 아래와 같이 실행 시킬 수 있다

```go
func main() {
	user := User{Name: "a", Email: "a@b.c", Age: 100}
	user2 := User{Name: "d", Email: "e@f.g", Age: 101}
 	
  // Parse -> ParseFiles
	tmpl, err := template.New("Tmpl1").ParseFiles("./templates/tmpl1.tmpl")
	if err != nil {
		panic(err)
	}
  
  // Execute -> ExecuteTemplate
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
}
```

```
Name: {{.Name}}
Email: {{.Email}}
Age : {{.Age}}

// tmpl1.tmpl
```