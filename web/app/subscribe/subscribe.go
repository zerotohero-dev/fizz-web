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
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile := session.Get("profile")

		if profile == nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}

		_, subscribed := (profile.(map[string]interface{}))["subscribed"]
		if subscribed {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		ctx.HTML(http.StatusOK, "subscribe.html", profile)
	}
}
