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

package logout

import (
	"github.com/gin-gonic/gin"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
	"net/url"
	"os"
)

// Handler for our logout.
func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logoutUrl, err := url.Parse("https://" + os.Getenv("FIZZ_WEB_AUTH0_DOMAIN") + "/v2/logout")
		if err != nil {
			log.Err("Internal server error: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Internal server error.")
			return
		}

		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
		if err != nil {
			log.Err("Internal server error: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Internal server error.")
			return
		}
		// Due to TLS termination at ALB, ctx.Request.TLS is `false` for production.
		// Check the host string instead.
		if returnTo.String() == "http://fizzbuzz.pro" ||
			returnTo.String() == "http://www.fizzbuzz.pro" {
			returnTo, err = url.Parse("https://fizzbuzz.pro")
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", os.Getenv("FIZZ_WEB_AUTH0_CLIENT_ID"))
		logoutUrl.RawQuery = parameters.Encode()

		toRedirect := logoutUrl.String()
		ctx.Redirect(http.StatusTemporaryRedirect, toRedirect)
	}
}