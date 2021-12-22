/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package middleware

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// IsSubscribed checks Gumroad API for a matching subscription for the user.
func IsSubscribed(ctx *gin.Context) {

	// TODO: use a common struct and validate all of these env vars during app initialization.

	gumroadProductId := os.Getenv("FIZZ_WEB_GUMROAD_PRODUCT_ID")
	if gumroadProductId == "" {
		log.Fatal("gumroadProductId not set.")
		return
	}

	gumroadAccessToken := os.Getenv("FIZZ_WEB_GUMROAD_ACCESS_TOKEN")
	if gumroadAccessToken == "" {
		log.Fatal("gumroadAccessToken not set.")
		return
	}

	// Cannot check subscription if the user does not have an account.
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Parse user’s profile
	p := sessions.Default(ctx).Get("profile")
	profile, ok := p.(map[string]interface{})
	if !ok {
		log.Fatal("Effed up session!")
		return
	}

	// If user is already subscribed, send them to where they need to go.
	_, subscribed := profile["subscribed"]
	if subscribed {
		ctx.Next()
		return
	}

	email, ok := profile["email"].(string)
	if !ok {
		log.Err("auth0 not able to fetch profile email")
		ctx.Next()
		return
	}
	if email == "" {
		log.Err("auth0 profile email appears to be blank")
		ctx.Next()
		return
	}

	// TODO: maybe add an exponential backoff if gumroad api fails.
	apiUrl := "https://api.gumroad.com/v2/products/" +
		gumroadProductId +
		"/subscribers?access_token=" +
		gumroadAccessToken + "&email=" +
		url.QueryEscape(email)

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Err("Gumroad API failure: %s", err.Error())
		ctx.Next()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err("Gumroad API parse failure: %s", err.Error())
		ctx.Next()
		return
	}

	var gr GumroadResponse

	err = json.Unmarshal(body, &gr)
	if err != nil {
		log.Err("Gumroad unmarshal error: %s", err.Error())
		ctx.Next()
		return
	}

	// If Gumroad does not succeed, it is not the user’s problem.
	// Just log the error, and let the user in this time.
	if !gr.Success {
		log.Err("Gumroad unsuccessful response")
		ctx.Next()
		return
	}

	// No subscriber found. Let the user subscribe.
	if len(gr.Subscribers) == 0 {
		ctx.Redirect(http.StatusSeeOther, "/subscribe")
		return
	}

	// Mark user as “subscribed” so that we don’t do redundant Gumroad API
	// lookups.
	profile["subscribed"] = true
	sessions.Default(ctx).Set("profile", profile)
	ctx.Next()
}
