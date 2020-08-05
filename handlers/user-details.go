package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parikshitg/goauth2/models"
	s "github.com/parikshitg/goauth2/sessions"
)

// UserDetailsHandler Handler
func UserDetailsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	data := make(map[string]interface{})
	data["title"] = "User Details"

	page, err := template.ParseFiles("templates/user-details.html", "templates/footer.html", "templates/header2.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	session, _ := s.Store.Get(r, "auth-cookie")
	useremail, _ := session.Values["Useremail"]
	user, _ := models.ExistingUser(useremail.(string))
	data["User"] = user.Name

	userDetail, _ := models.UserDetails(id)
	data["UserDetails"] = userDetail

	var mett Meta
	err = json.Unmarshal([]byte(userDetail.Meta), &mett)
	if err != nil {
		log.Println("error : ", err)
	}

	md, _ := mett.Github.(map[string]interface{})
	log.Println("git name : ", md["name"])
	log.Println("git Email : ", md["email"])
	log.Println("git Url : ", md["url"])
	log.Println("git Username : ", md["login"])

	data["GitName"] = md["name"]
	data["GitEmail"] = md["email"]
	data["GitUrl"] = md["url"]
	data["GitUsername"] = md["login"]

	mdd, _ := mett.Linkedin.(map[string]interface{})
	log.Println("linkedin First Name : ", mdd["localizedFirstName"])
	log.Println("linkedin Last Name : ", mdd["localizedLastName"])
	log.Println("linkedin email : ", mdd["email"])

	data["LFName"] = mdd["localizedFirstName"]
	data["LLName"] = mdd["localizedLastName"]
	data["LEmail"] = mdd["email"]

	log.Println("twitter : ", mett.Twitter)

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
