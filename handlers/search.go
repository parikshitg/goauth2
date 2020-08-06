package handlers

import (
	"html/template"
	"log"
	"net/http"

	"goauth2/models"
	s "goauth2/sessions"
)

// Search Handler
func Search(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Search User"

	page, err := template.ParseFiles("templates/search.html", "templates/footer.html", "templates/header2.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	session, _ := s.Store.Get(r, "auth-cookie")
	useremail, _ := session.Values["Useremail"]
	user, _ := models.ExistingUser(useremail.(string))
	data["User"] = user.Name

	if r.Method == http.MethodPost {

		email := r.FormValue("email")

		userDetails, ok := models.ExistingUser(email)
		if !ok {
			data["Message"] = "User Does Not Exists !!"
		} else {
			data["Message"] = "User Found"
			data["Details"] = userDetails
		}
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
