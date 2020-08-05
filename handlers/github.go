package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/models"
	"github.com/parikshitg/goauth2/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Github application client id and secret key
const clientID = "7ca5db25b59a728eca66"
const clientSecret = "fe0bffeafe45b057036d0ea9ffb619787eb38583"

var config = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Endpoint:     github.Endpoint,
	RedirectURL:  "http://localhost:8080" + "/github/callback",
	Scopes:       []string{"user"},
}

type GitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Url      string `json:"url"`
	Username string `json:"login"`
}

type Meta struct {
	Github   interface{} `json:"github"`
	Linkedin interface{} `json:"linkedin"`
	Twitter  interface{} `json:"twitter"`
}

// Github Login handler
func GithubLogin(w http.ResponseWriter, r *http.Request) {

	state := sessions.NewOauth2State()
	log.Println("state:", state)
	url := config.AuthCodeURL(state)
	log.Println("url:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Github Callback Handler
func GithubCallback(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	log.Println("callback state:", state)

	if sessions.VerifyOauth2State(state) {
		log.Printf("Invalid oauth state: %s\n", state)
		return
	}

	code := r.FormValue("code")
	log.Println("code:", code)
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Code exchange failed with '%s'\n", err)
		return
	}

	client := config.Client(oauth2.NoContext, token)
	response, err := client.Get("https://api.github.com/user")

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	// Unmarshal contents into type GitUser
	var gituser GitUser
	err = json.Unmarshal(contents, &gituser)
	if err != nil {
		log.Println("unmarshal error : ")
		return
	}

	gitMetaData := &GitUser{Name: gituser.Name, Email: gituser.Email, Username: gituser.Username, Url: gituser.Url}

	dbuser, present := models.ExistingUser(gituser.Email)
	if present {

		var met Meta
		err := json.Unmarshal([]byte(dbuser.Meta), &met)
		if err != nil {
			log.Println("unmarshal Error : ", err)
			return
		}

		if met.Github == "" {

			m := make(map[string]interface{})
			m["Github"] = gitMetaData
			m["Linkedin"] = met.Linkedin
			m["Twitter"] = met.Twitter

			// Marshal map to store as a string into database
			v, err := json.Marshal(m)
			if err != nil {
				log.Println("Marshal error: ", err)
				return
			}

			models.Db.Debug().Table("users").Where("email = ?", gituser.Email).Update("meta", v)
		}
	} else {

		mm := make(map[string]interface{})
		mm["Github"] = gitMetaData
		mm["Linkedin"] = ""
		mm["Twitter"] = ""
		// Marshal map to store as a string into database
		v, err := json.Marshal(mm)
		if err != nil {
			log.Println("Marshal error: ", err)
			return
		}

		user := &models.User{Name: gituser.Name, Email: gituser.Email, Meta: string(v)}
		// else create a new user
		models.Db.Debug().Create(&user)
	}

	// setting up a session (Login)
	session, err := sessions.Store.Get(r, "auth-cookie")
	if err != nil {
		log.Println("Session Error:", err)
		return
	}
	session.Values["Useremail"] = gituser.Email
	session.Save(r, w)

	http.Redirect(w, r, "/user/all", http.StatusSeeOther)
}
