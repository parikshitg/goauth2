package handlers

import (
	"fmt"
	"net/http"

	"github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/dghubble/sessions"
)

var Oauth1Config = &oauth1.Config{
	ConsumerKey:    "5wQgXq0udwjyHMCiOc5W5F2E2",
	ConsumerSecret: "DvTmXUbz9wN9fLfgR2XJBLvLvAPxG8cL2M3TTXFGKqtcYJaC32",
	CallbackURL:    "http://localhost:8080/twitter/callback",
	Endpoint:       twitterOAuth1.AuthorizeEndpoint,
}

const (
	sessionName     = "example-twtter-app"
	sessionSecret   = "example cookie signing secret"
	sessionUserKey  = "twitterID"
	sessionUsername = "twitterUsername"
)

// sessionStore encodes and decodes session data stored in signed cookies
var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// issueSession issues a cookie session after successful Twitter login
func IssueSession() http.Handler {
	fmt.Println("Executed issues session")
	fn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside handler ")
		ctx := req.Context()
		twitterUser, err := twitter.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("name : ", twitterUser.ScreenName)
		// 2. Implement a success handler to issue some form of session
		session := sessionStore.New(sessionName)
		session.Values[sessionUserKey] = twitterUser.ID
		session.Values[sessionUsername] = twitterUser.ScreenName
		session.Save(w)
		http.Redirect(w, req, "/register", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}
