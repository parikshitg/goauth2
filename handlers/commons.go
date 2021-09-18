package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/parikshitg/goauth2/sessions"
)

type Meta struct {
	Github   interface{} `json:"github"`
	Linkedin interface{} `json:"linkedin"`
	Twitter  interface{} `json:"twitter"`
}

// Create map to insert meta into database
func MakeMetaMap(git, link, twit interface{}) []byte {

	m2 := make(map[string]interface{})
	m2["Github"] = git
	m2["Linkedin"] = link
	m2["Twitter"] = twit

	val, err := json.Marshal(m2)
	if err != nil {
		log.Println("Marshal error: ", err)
		return nil
	}

	return val
}

// Set Login Session
func SetSession(email string, w http.ResponseWriter, r *http.Request) {

	session, err := sessions.Store.Get(r, "auth-cookie")
	if err != nil {
		log.Println("Session Error:", err)
		return
	}
	session.Values["Useremail"] = email
	session.Save(r, w)
}
