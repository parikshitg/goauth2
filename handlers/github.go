package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"github.com/parikshitg/goauth2/sessions"
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
	Name string `json:"name"`
	Email string `json:"email"`
}

// Github Login handler
func GithubLogin(w http.ResponseWriter, r *http.Request) {

	state := sessions.NewOauth2State()
	log.Println("state:", state)
	url := config.AuthCodeURL(state)
	log.Println("url:", url)
	http.Redirect(w,r, url,http.StatusTemporaryRedirect)
}

// Callback Handler
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
	
	var gituser GitUser
	err = json.Unmarshal(contents, &gituser)
	if err != nil {
		log.Println("unmarshal error : ")
		return
	}
	log.Println("gituser  : ", gituser)

	http.Redirect(w,r ,"/dashboard", http.StatusSeeOther)
}
