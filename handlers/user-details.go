package handlers

import (
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

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
