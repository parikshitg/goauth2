package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	t "github.com/dghubble/go-twitter/twitter"
	oauth1Login "github.com/dghubble/gologin/oauth1"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/parikshitg/goauth2/models"
)

var Oauth1Config = &oauth1.Config{
	ConsumerKey:    "5wQgXq0udwjyHMCiOc5W5F2E2",
	ConsumerSecret: "DvTmXUbz9wN9fLfgR2XJBLvLvAPxG8cL2M3TTXFGKqtcYJaC32",
	CallbackURL:    "http://localhost:8080/twitter/callback",
	Endpoint:       twitterOAuth1.AuthorizeEndpoint,
}

func IssueSession() http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		accessToken, accessSecret, err := oauth1Login.AccessTokenFromContext(ctx)
		if err != nil {
			fmt.Println("err")
			return
		}

		httpClient := Oauth1Config.Client(ctx, oauth1.NewToken(accessToken, accessSecret))
		twitterClient := t.NewClient(httpClient)
		accountVerifyParams := &t.AccountVerifyParams{
			IncludeEntities: t.Bool(true),
			SkipStatus:      t.Bool(false),
			IncludeEmail:    t.Bool(true),
		}

		user, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)

		if err != nil {
			fmt.Println("Errr")
			return
		}

		// fmt.Println("username : ", user.ScreenName)
		// fmt.Println("name : ", user.Name)
		// fmt.Println("email : ", user.Email)
		// fmt.Println("name : ", user.ID)

		tu := &TwitUser{Username: user.ScreenName, Name: user.Name, Email: user.Email, Location: user.Location}

		dbuser, present := models.ExistingUser(user.Email)
		if present {

			var met Meta
			err := json.Unmarshal([]byte(dbuser.Meta), &met)
			if err != nil {
				log.Println("unmarshal Error : ", err)
				return
			}

			if met.Twitter == "" {

				val := MakeMetaMap(met.Github, met.Linkedin, tu)
				models.Db.Debug().Table("users").Where("email = ?", user.Email).Update("meta", val)
			}
		} else {

			val := MakeMetaMap("", "", tu)
			user := &models.User{Name: user.Name, Email: user.Email, Meta: string(val)}
			models.Db.Debug().Create(&user)
		}

		SetSession(user.Email, w, r)
		http.Redirect(w, r, "/user/all", http.StatusSeeOther)
	}

	return http.HandlerFunc(fn)
}

type TwitUser struct {
	Username string
	Name     string
	Email    string
	Location string
}
