package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

// Linkedin application client id and secret key
const linkedinClientID = "868h74ftndcy4z"
const linkedinClientSecret = "iDl6UnlhKrkiM1sM"

var lconfig = &oauth2.Config{
	ClientID:     linkedinClientID,
	ClientSecret: linkedinClientSecret,
	Endpoint:     linkedin.Endpoint,
	RedirectURL:  "http://localhost:8080/linkedin/callback",
	Scopes:       []string{"r_emailaddress", "r_liteprofile"},
}

// Linkedin Login Handler
func LinkedinLogin(w http.ResponseWriter, r *http.Request) {
	state := sessions.NewOauth2State()
	log.Println("state:", state)
	url := lconfig.AuthCodeURL(state)
	log.Println("url:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Linkedin callback handler
func LinkedinCallback(w http.ResponseWriter, r *http.Request) {

	log.Println("error:", r.FormValue("error"))
	log.Println("error_description:", r.FormValue("error_description"))
	state := r.FormValue("state")
	log.Println("callback state:", state)

	if sessions.VerifyOauth2State(state) {
		log.Printf("Invalid oauth state: %s\n", state)
		return
	}

	code := r.FormValue("code")
	log.Println("code:", code)
	token, err := lconfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Code exchange failed with '%s'\n", err)
		return
	}

	client := lconfig.Client(oauth2.NoContext, token)

	// Api for email address
	response, err := client.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	var p map[string]interface{}
	json.Unmarshal(contents, &p)
	email := p["elements"].([]interface{})[0].(map[string]interface{})["handle~"].(map[string]interface{})["emailAddress"]

	log.Println("email : ", email)

	// Api for name
	response2, err := client.Get("https://api.linkedin.com/v2/me")
	defer response.Body.Close()
	contents2, err := ioutil.ReadAll(response2.Body)
	fmt.Fprintf(w, "%s\n", contents2)

	var luser LinkedinUser
	err = json.Unmarshal(contents2, &luser)
	if err != nil {
		log.Println("unmarshal error : ")
		return
	}
	log.Println("luser  : ", luser)

}

type LinkedinUser struct {
	Firstname string `json:"localizedFirstName"`
	Lastname  string `json:"localizedLastName"`
}
