package main

import (
	"html/template"
	"os"
)

type User struct {
    Name string
    Bio  string
    Age  int
    City string
		Attributes map[string]string
}

func main() {
    t, err := template.ParseFiles("hello.gohtml")
    if err != nil {
        panic(err)
    }

    users := []User{
			{
					Name: "John Doe",
					Bio: "Some bio",
					Age: 25,
					City: "San Francisco",
					Attributes: map[string]string{
						"Hobbies" : "Working out",
						"Skills" : "Go, JavaScript",
			},
			},
			{
					Name: "Jane Doe",
					Bio: "Another bio",
					Age: 28,
					City: "New York",
					Attributes: map[string]string{
						"Hobbies": "Reading",
						"Skills": "Python, Ruby",
					},
			},
			{
					Name: "Michael Smith",	
					Bio: "Yet another bio",
					Age: 30,
					City: "Los Angeles",
					Attributes: map[string]string{
						"Hobbies": "Swimming",
						"Skills": "Java, C",
					},
			},
			}
		
		data := map[string]interface{}{
			"Users": users,
			"Header": "List of users",
		}

    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}
