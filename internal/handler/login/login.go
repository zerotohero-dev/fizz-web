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

package login

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-web/internal/authenticator"
	"net/http"
)

func generateSecureRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

// Handler for auth0 login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateSecureRandomState()
		if err != nil {
			log.Err("Error in authenticator: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Authentication error.")
			return
		}

		// Save the state into the session, to compare it when we receive it
		// on the callback endpoint.
		session := sessions.Default(ctx)
		session.Set("state", state)
		err = session.Save()
		if err != nil {
			log.Err("Error in setting session state: %s", err.Error())
			ctx.String(http.StatusInternalServerError, "Session error")
			return
		}

		// Redirect to provider login:
		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}
