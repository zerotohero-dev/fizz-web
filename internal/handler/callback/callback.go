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

package callback

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-web/internal/authenticator"
	"net/http"
)

// Handler for auth0 callback.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}

		// Exchange authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(
				http.StatusUnauthorized,
				"Failed to convert an authorization code into a token.",
			)
			return
		}

		// verify token, an get id token.
		idToken, err := auth.VerifyIdToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		// Parse id token’s claims into profile.
		var profile map[string]interface{}
		err = idToken.Claims(&profile)
		if err != nil {
			log.Err("Error claiming profile: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Failed to claim profile.")
			return
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		err = session.Save()
		if err != nil {
			log.Err("Failed to save session: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Failed to save session.")
			return
		}

		// Redirect back home.
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
