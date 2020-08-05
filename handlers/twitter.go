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
	"github.com/parikshitg/goauth2/conf"
	"github.com/parikshitg/goauth2/models"
)

var Oauth1Config = &oauth1.Config{
	ConsumerKey:    conf.TwitterConsumerKey,
	ConsumerSecret: conf.TwitterConsumerSecret,
	CallbackURL:    "http://localhost:8080/twitter/callback",
	Endpoint:       twitterOAuth1.AuthorizeEndpoint,
}

type TwitUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
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
