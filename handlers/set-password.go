package handlers

import (
	"html/template"
	"log"
	"net/http"

	"goauth2/models"
	s "goauth2/sessions"
)

// SetPassword Handler
func SetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Set Password"

	page, err := template.ParseFiles("templates/set-password.html", "templates/footer.html", "templates/header2.html")
	if err != nil {
		log.Println("ParseFiles: ", err)
		return
	}

	session, _ := s.Store.Get(r, "auth-cookie")
	useremail, _ := session.Values["Useremail"]
	user, _ := models.ExistingUser(useremail.(string))
	data["User"] = user.Name

	if user.Password != "" {
		data["Pass"] = true

		if r.Method == http.MethodPost {

			password := r.FormValue("password")
			password2 := r.FormValue("password2")
			password3 := r.FormValue("password3")

			if password == "" || password2 == "" || password3 == "" {

				data["Flash"] = "Fields can not be Empty !!"

				err = page.Execute(w, data)
				if err != nil {
					log.Fatal("Execute:", err)
				}
				return
			}

			if password2 != password3 {

				data["Flash"] = "New Passwords Did not match !!"

				err = page.Execute(w, data)
				if err != nil {
					log.Fatal("Execute:", err)
				}
				return
			}

			if password != user.Password {

				data["Flash"] = "Invalid Current Password !!"

				err = page.Execute(w, data)
				if err != nil {
					log.Fatal("Execute:", err)
				}
				return
			}

			models.SetNewPass(useremail.(string), password2)
			data["Flash"] = "Password Updated."
		}

	} else {

		if r.Method == http.MethodPost {

			password2 := r.FormValue("password2")
			password3 := r.FormValue("password3")

			if password2 == "" || password3 == "" {

				data["Flash"] = "Fields can not be Empty !!"

				err = page.Execute(w, data)
				if err != nil {
					log.Fatal("Execute:", err)
				}
				return
			}

			if password2 != password3 {

				data["Flash"] = "New Passwords Did not match !!"

				err = page.Execute(w, data)
				if err != nil {
					log.Fatal("Execute:", err)
				}
				return
			}

			models.SetNewPass(useremail.(string), password2)
			data["Flash"] = "Password Updated."
		}
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
