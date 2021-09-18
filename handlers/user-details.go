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

	userDetail, ok := models.UserDetails(id)
	if !ok {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	data["UserDetails"] = userDetail

	var mett Meta
	err = json.Unmarshal([]byte(userDetail.Meta), &mett)
	if err != nil {
		log.Println("error : ", err)
	}

	md, _ := mett.Github.(map[string]interface{})
	data["GitName"] = md["name"]
	data["GitEmail"] = md["email"]
	data["GitUrl"] = md["url"]
	data["GitUsername"] = md["login"]

	md2, _ := mett.Linkedin.(map[string]interface{})
	data["LFName"] = md2["localizedFirstName"]
	data["LLName"] = md2["localizedLastName"]
	data["LEmail"] = md2["email"]

	md3, _ := mett.Twitter.(map[string]interface{})
	data["TName"] = md3["name"]
	data["TUsername"] = md3["username"]
	data["TLocation"] = md3["location"]

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
