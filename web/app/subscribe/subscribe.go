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

package subscribe

import (
	"github.com/gin-contrib/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	// You need to be logged in to subscribe.
	// Normally middleware.isAuthenticated already takes care of this.
	// This check here is just defensive coding.
	if profile == nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// If user is already subscribed, send them home.
	_, subscribed := (profile.(map[string]interface{}))["subscribed"]
	if subscribed {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	// (based on the session state) The user does not appear to be subscribed;
	// render the subscription form.
	ctx.HTML(http.StatusOK, "subscribe.html", profile)
}
