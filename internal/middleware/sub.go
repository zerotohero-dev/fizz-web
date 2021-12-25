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

func IsNotSubscribed(ctx *gin.Context) {
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

	_, subscribed := profile["subscribed"]
	if subscribed {
		// Already subscribed, no need to show the subscription page again.
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	email, ok := profile["email"].(string)
	if !ok || email == "" {
		log.Err("auth0 not able to fetch profile email")
		// Unable to fetch email; still allow user to consume the content:
		// We can manually fix things later.
		ctx.Next()
		return
	}

	apiUrl := "https://api.gumroad.com/v2/products/" +
		gumroadProductId +
		"/subscribers?access_token=" +
		gumroadAccessToken + "&email=" +
		url.QueryEscape(email)

	resp, err := http.Get(apiUrl)
	if err != nil {
		// Gumroad API failure is not the user’s problem.
		// Let them consume the content until we figure out what’s wrong.
		log.Err("gumroad api failure: %s", err.Error())
		ctx.Next()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Again, Gumroad API failure is not the user’s problem.
		// Be optimistic, and allow access to content until things are sorted out.
		log.Err("gumroad api failure: %s", err.Error())
		ctx.Next()
		return
	}

	var gr GumroadResponse
	err = json.Unmarshal(body, &gr)
	if err != nil {
		// Same. Your integration error, your problem: Don’t punish the users for that.
		log.Err("Gumroad unmarshal error: %s", err.Error())
		ctx.Next()
		return
	}

	if !gr.Success {
		// Allow the user to proceed regardless.
		log.Err("Gumroad response failure: %s", err.Error())
		ctx.Next()
		return
	}

	// We are certain that the user is not subscribed.
	// Allow to render the subscription form.
	if len(gr.Subscribers) == 0 {
		ctx.Next()
		return
	}

	// User is subscribed, but profile is stale; update it.
	profile["subscribed"] = true
	session := sessions.Default(ctx)
	session.Set("profile", profile)
	err = session.Save()
	if err != nil {
		log.Err("Failed to save session: %s", err.Error())
	}

	// If subscribed, yeet the user to the default page.
	ctx.Redirect(http.StatusSeeOther, "/")
}

func IsSubscribed(ctx *gin.Context) {
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

	_, subscribed := profile["subscribed"]
	if subscribed {
		ctx.Next()
		return
	}

	email, ok := profile["email"].(string)
	if !ok || email == "" {
		log.Err("auth0 not able to fetch profile email")
		// Unable to fetch email; still allow user to consume the content:
		// We can manually fix things later.
		ctx.Next()
		return
	}

	apiUrl := "https://api.gumroad.com/v2/products/" +
		gumroadProductId +
		"/subscribers?access_token=" +
		gumroadAccessToken + "&email=" +
		url.QueryEscape(email)

	resp, err := http.Get(apiUrl)
	if err != nil {
		// Gumroad API failure is not the user’s problem.
		// Let them consume the content until we figure out what’s wrong.
		log.Err("gumroad api failure: %s", err.Error())
		ctx.Next()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Again, Gumroad API failure is not the user’s problem.
		// Be optimistic, and allow access to content until things are sorted out.
		log.Err("gumroad api failure: %s", err.Error())
		ctx.Next()
		return
	}

	var gr GumroadResponse
	err = json.Unmarshal(body, &gr)
	if err != nil {
		// Same. Your integration error, your problem: Don’t punish the users for that.
		log.Err("Gumroad unmarshal error: %s", err.Error())
		ctx.Next()
		return
	}

	if !gr.Success {
		// Allow the user to proceed regardless.
		log.Err("Gumroad response failure: %s", err.Error())
		ctx.Next()
		return
	}

	// If no subscribers, then we are mostly certain that the user is not
	// subscribed yet. Redirect them to the subscription page.
	if len(gr.Subscribers) == 0 {
		ctx.Redirect(http.StatusSeeOther, "/subscribe")
		return
	}

	profile["subscribed"] = true
	session := sessions.Default(ctx)
	session.Set("profile", profile)
	err = session.Save()
	if err != nil {
		log.Err("Failed to save session: %s", err.Error())
	}
	ctx.Next()
}
