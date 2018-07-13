package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type User struct {
	ID      string
	Name    string
	Email   string
	Phone   string
	Address string
	City    string
}

type PageVariables struct {
	Date string
	Time string
}

var users []User

func main() {
	initializeData()

	s := "https://example.com/?id=1"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	param, _ := url.ParseQuery(u.RawQuery)
	id := param["id"][0]
	fmt.Println("ID: ", id)

	// http.HandleFunc("/user/{id}", UserPage)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", UserPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8000", nil)
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsID := params["id"]
	var UserVars User

	// queryURL := r.URL.Query()
	// id := queryURL.Get("id")

	for _, user := range users {
		if paramsID == user.ID {
			UserVars = User{
				ID:      user.ID,
				Name:    user.Name,
				Email:   user.Email,
				Phone:   user.Phone,
				Address: user.Address,
				City:    user.City,
			}
		}
	}

	t, err := template.ParseFiles("userpage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, UserVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// var id, name, email, phone, address, city string

	// queryURL := r.URL.Query()
	// id = queryURL.Get("id")

	// // fmt.Println("URL param id is: " + string(id))

	// for _, user := range users {
	// 	if id == user.ID {
	// 		name = user.Name
	// 		email = user.Email
	// 		phone = user.Phone
	// 		address = user.Address
	// 		city = user.City
	// 	}
	// }

	// io.WriteString(w, "Name: "+name)
	// io.WriteString(w, "\nEmail: "+email)
	// io.WriteString(w, "\nPhone: "+phone)
	// io.WriteString(w, "\nAddress: "+address)
	// io.WriteString(w, "\nCity: "+city)
}

func initializeData() {
	users = append(users, User{ID: "1", Name: "Rhoma", Email: "rhoma@gmail.com", Phone: "081123", Address: "Sidoarjo", City: "Sidoarjo"})
	users = append(users, User{ID: "2", Name: "Cahyanti", Email: "cahyanti@gmail.com", Phone: "081123", Address: "Sidoarjo", City: "Sidoarjo"})
	users = append(users, User{ID: "3", Name: "John", Email: "john@gmail.com", Phone: "081123", Address: "Brooklyn", City: "New York"})
	users = append(users, User{ID: "4", Name: "Doe", Email: "doe@gmail.com", Phone: "081123", Address: "Brooklyn", City: "New York"})
	users = append(users, User{ID: "5", Name: "Harry", Email: "harry@gmail.com", Phone: "081123", Address: "Privet Drive", City: "London"})
	users = append(users, User{ID: "6", Name: "Potter", Email: "potter@gmail.com", Phone: "081123", Address: "Privet Drive", City: "London"})
}
